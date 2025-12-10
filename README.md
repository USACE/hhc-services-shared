# HH&C Shared Database and API

## API Environment Variables

| Environment Variable Name | Local Development | AWS Defined | API Config Required | API Config Default | AWS Value                    |
| :------------------------ | :---------------: | :---------: | :-----------------: | :----------------- | :--------------------------- |
| AWS_ACCESS_KEY_ID         |         X         |             |                     |                    |                              |
| AWS_SECRET_ACCESS_KEY     |         X         |             |                     |                    |                              |
| AWS_DEFAULT_REGION        |         X         |      X      |                     |                    | defined in IaC               |
| AWS_DISABLE_SSL           |         X         |             |                     |                    |                              |
| AWS_S3_FORCE_PATH_STYLE   |         X         |             |                     |                    |                              |
| AUTH_ENVIRONMENT          |         X         |      X      |                     |                    | defined in IaC               |
| APPLICATION_KEY           |         X         |      X      |                     |                    | defined in IaC               |
| MINIO_ENDPOINT_URL        |         X         |             |                     |                    |                              |
| PGUSER                    |         X         |      X      |          X          |                    | hhc_shared_user              |
| PGPASSWORD                |         X         |      X      |          X          |                    | defined in IaC               |
| PGDATABASE                |         X         |      X      |          X          |                    | postgres                     |
| PGHOST                    |         X         |      X      |          X          |                    | defined in IaC               |
| PGSSLMODE                 |         X         |      X      |          X          | require            | require                      |
| PGX_POOL_MAXCONNS         |         X         |             |                     | 10                 |                              |
| PGX_POOL_MINCONNS         |         X         |             |                     | 5                  |                              |
| PGX_POOL_MAXCONN_IDLETIME |         X         |             |                     | 30m                |                              |
| S3_BUCKET                 |         X         |      X      |          X          |                    | hhc-shared-***ENVIRONMENT*** |
| S3_DEFAULT_INDEX          |         X         |             |                     | index.html         |                              |
| S3_PREFIX_STATIC          |         X         |      X      |                     | /                  | /shared/ui                   |
| API_PORT                  |         X         |             |                     | 8080               |                              |
| API_LOG                   |         X         |             |                     | false              |                              |


## Managing spatial data with the PostGIS extension
The following setup steps for the PostGIS extension is not available through automation processes.  PostGIS extension setup requires rds_superuser privileges.  Amazon RDS documentation [here](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.PostGIS.html) is a guide to help with initial setup in addition to the following.

### PostGIS Extension Setup

1. Make a connection to the database
   1. Connection to the database needs to be a user with elevated privileges
   2. The typical user is `postgres`
   3. Acquire this user's password from the `AWS Secrets Manager`
2. Create a `ROLE` to manage the PostGIS extension and `GRANT` `rds_superuser` to `gis_admin`

   ```sql
    CREATE ROLE gis_admin;
    GRANT rds_superuser TO gis_admin;
    ```
3. Set role to `gis_admin` and create extension(s)
   1. This will add three additional schemas, tiger, tiger_data, and topology
   2. The `public` schema will have the table `spatial_ref_sys` created

    ```sql
    SET ROLE gis_admin;

    CREATE EXTENSION postgis;
    CREATE EXTENSION fuzzystrmatch;
    CREATE EXTENSION postgis_tiger_geocoder;
    CREATE EXTENSION postgis_topology;
    CREATE EXTENSION address_standardizer_data_us; -- optional
    ```
4. Verify the extensions and their owners

    ```sql
    SELECT
        n.nspname AS "Name"
        , pg_catalog.pg_get_userbyid(n.nspowner) AS "Owner"
    FROM
        pg_catalog.pg_namespace n
    WHERE
        n.nspname !~ '^pg_'
        AND n.nspname <> 'information_schema'
    ORDER BY
        1;
    ```

    Expected Result

    | Name       | Owner    |
    | :--------- | :------- |
    | public     | postgres |
    | tiger      | rdsadmin |
    | tiger_data | rdsadmin |
    | topology   | rdsadmin |

5. Transfer ownership of the extension schemas to the `gis_admin` role

    ```sql
    ALTER SCHEMA tiger OWNER TO gis_admin;
    ALTER SCHEMA tiger_data OWNER TO gis_admin;
    ALTER SCHEMA topology OWNER TO gis_admin;
    ```
6. Verify the extensions and their owners once again

    ```sql
    SELECT
        n.nspname AS "Name"
        , pg_catalog.pg_get_userbyid(n.nspowner) AS "Owner"
    FROM
        pg_catalog.pg_namespace n
    WHERE
        n.nspname !~ '^pg_'
        AND n.nspname <> 'information_schema'
    ORDER BY
        1;
    ```

    Expected Result
    | Name       | Owner     |
    | :--------- | :-------- |
    | public     | postgres  |
    | tiger      | gis_admin |
    | tiger_data | gis_admin |
    | topology   | gis_admin |

7. Transfer ownership of the PostGIS tables
   1. Create a function to alter permissions
   
   ```sql
   CREATE FUNCTION exec(text) returns text language plpgsql volatile AS $f$ BEGIN EXECUTE $1; RETURN $1; END; $f$;
   ```

   2. Run the query that runs the function that alters permissions

    ```sql
    SELECT
        exec ('ALTER TABLE ' || quote_ident(s.nspname) || '.' || quote_ident(s.relname) || ' OWNER TO gis_admin;')
    FROM (
        SELECT
            nspname
            , relname
        FROM
            pg_class c
            JOIN pg_namespace n ON (c.relnamespace = n.oid)
        WHERE
            nspname IN ('tiger' , 'topology')
            AND relkind IN ('r' , 'S' , 'v')
        ORDER BY
            relkind = 'S') s;
    ```

8. Testing the extension
   1. Set the search path to avoid needing to specify the schema name

    ```sql
    SET search_path=public,tiger;
    ```

9. Test the `tiger` schema with the following

    ```sql
    SELECT address, streetname, streettypeabbrev, zip
    FROM normalize_address('1 Devonshire Place, Boston, MA 02109') AS na;
    ```

    Expected Result

    | address | streetname | streettypeabbrev | zip   |
    | :------ | :--------- | :--------------- | :---- |
    | 1       | Devonshire | Pl               | 02109 |

### PostGIS Extension Versions

To get a list of versions use the following command

```sql
SELECT * FROM pg_available_extension_versions WHERE name='postgis';
```

### PostGIS Extension Upgrade

Check for available PostGIS extension version updates by running the following

```sql
SELECT postGIS_extensions_upgrade();
```

## PostgreSQL Extension(s)

In addition to the PostGIS extensions, there is a requirement for the `uuid-ossp` extension.  Many schema tables use the `uuid_generate_v4()` function as a data type.  This extension needs to associated with the `public` schema to automatically be available to all other schemas.

Create the extension for `uuid-ossp`

```sql
CREATE EXTENSION "uuid-ossp" WITH SCHEMA public;
```

## Flyway Migration Placeholders

Flyway comes with support for placeholder replacement in:

- SQL migrations
- Script migrations

[migration-placeholder](https://documentation.red-gate.com/flyway/flyway-concepts/migrations/migration-placeholders)

[placeholder-namespace](https://documentation.red-gate.com/flyway/reference/configuration/flyway-namespace/flyway-placeholders-namespace)


## Instructions Building/Updating `geometry.tar.gz`

Due to file size limitations (>100MB) for GitHub repositories, large geometry files supporting the Flyway migration are offloaded to the repository's release assets.  Assets are collected in a `tar.gz` file and associated with a tagged release.  Developing locally with needs for large files supporting new migrations will require a download of the asset making sure you have the most up-to-date set of files.  Make sure to include any new files and update the asset in the release.

### Create the Tar gzip

Tar gzip all the zip files into `geometry.tar.gz`.  These are the supporting files for migration that are too large (>100MB) for the repository.  Building the tar gzip and uploading to the repository as taged release asset has no size limitation.

The tar gzip file is a collection of all the zip files used in the migration.  Typically, a migration file uses a single zip file.  The following command is an example how the tar gzip file could be created:

```bash
> tar -czvf geometry.tar.gz /path/to/*.zip
```

### Generate District Geometry: The Python Script

The Python script `master_data_geom.py` reads geojson from ARC GIS services, defining District boundaries, and generates a SQL file that can be used to INSERT that data into a Postgres database.  There are four URLs defining different missions:

    - Civil Works
    - Regulatory
    - FUDS
    - Military

Not every District has these boundaries defined.  SQL files are generated if that District has geojson defined for a particular mission, and that file is written to a zip file.  The result of running this script is a collection of SQL files, one for each District, each having their respective boundaries defined in the `mission_geom.zip` file.

## Update the UI
There is no role or CI/CD setup for this repo to update the UI to any CWBI environment.  This UI should not be updated that often, so not really worth the effort to have a role and go through the change management to get all the setup.<br><br>The following describes steps using `aws-adfs` login and some suggestions that could make uploads a bit easier.

### Build the ./dist
A bash script is available to make life a little easier to build the UI for the envrionment you want to update.  A `dotenv` file already has the configurations setup and do not need to be modified.  Use the `./build_mv.sh` script with its options to build or build/move for the respective CWBI environment.

```bash
Usage: ./build_mv.sh [-m] [-h] [dev|test|prod|clear|help]

No argument runs 'vite build'.  With argument runs 'vite build --mode arg',
where arg is dev, test, or prod. All commands result in moving the resulting
'./dist' files into ./_media/ui/.

  -m:    build and move to ./_media/shared/ui
  -h:    display this help message
  dev:   vite build using mode dev using env vars from .env.dev
  test:  vite build using mode test using env vars from .env.test
  prod:  vite build using mode prod using env vars from .env.prod
  clear: remove all files from './_media/shared/ui'

  There is a local variable PREFIX (./services/ui) that references the relative
  path to package.json. There is another local variable MEDIA (./_media/shared/ui) that
  references the relative path to media files.

  These two local variables should be modified to fit your repo's needs.
```

## Update the UI on AWS

### AWS Config Profile
Define a profile in the local file ~/.aws/config with suggested inputs below.  This allows for a simple and short aws-adfs login command:

```bash
aws-adfs login --profile PROFILE
```

Suggested Profile Configuration:

- region = REGION_NAME
- output = json
- adfs_config.ssl_verification = True
- adfs_config.role_arn = ROLE_ARN
- adfs_config.adfs_host = ADFS_HOST
- adfs_config.session_duration = SESSION_DURATION_SECONDS
- adfs_config.provider_id = PROVIDER_ID
- adfs_config.sspi = False
- adfs_config.duo_factor = None
- adfs_config.duo_device = None
- adfs_config.adfs_user = ADFS_USER


### Update AWS S3

A2W static content is served directly from an AWS S3 bucket, `a2w-dev`, `a2w-static-test`, and `a2w-static-prod`.  Each bucket requires files to be updated after UI changes and a merge into the `develop` branch.  Bucket prefix for Test and Prod are both `site/`.

*`a2w-dev` is currently setup with CI/CD

### Update Content, AWS CLI

- Make sure to ***first*** log into your account using `aws-adfs` on the command line.
- Make sure to replace `BUCKET_NAME` and `PREFIX`.
- Remove `--dryrun` when ready to execute on S3

Remove files from bucket/prefix/assets:
```bash
aws s3 --profile PROFILE_NAME rm s3://BUCKET_NAME/PREFIX/assets --recursive --dryrun
```

Copy files that have changed:
```bash
aws s3 --profile PROFILE_NAME sync ./dist s3://BUCKET_NAME/PREFIX --exclude "*.DS_Store" --dryrun
```

*`*.DS_Store` is for the Mac users to not copy those files to S3

All-in-one remove and update:
```bash
aws s3 --profile PROFILE_NAME rm s3://BUCKET_NAME/PREFIX/assets --recursive && aws s3 --profile PROFILE_NAME sync ./dist s3://BUCKET_NAME/PREFIX --exclude "*.DS_Store"
```
