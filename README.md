# Flyway Migration Placeholders

Flyway comes with support for placeholder replacement in:

- SQL migrations
- Script migrations

[migration-placeholder](https://documentation.red-gate.com/flyway/flyway-concepts/migrations/migration-placeholders)

[placeholder-namespace](https://documentation.red-gate.com/flyway/reference/configuration/flyway-namespace/flyway-placeholders-namespace)


# Instructions Building/Updating `geometry.tar.gz`

Due to file size limitations (>100MB) for GitHub repositories, large geometry files supporting the Flyway migration are offloaded to the repository's release assets.  Assets are collected in a `tar.gz` file and associated with a tagged release.  Developing locally with needs for large files supporting new migrations will require a download of the asset making sure you have the most up-to-date set of files.  Make sure to include any new files and update the asset in the release.

# Create the Tar gzip

Tar gzip all the zip files into `geometry.tar.gz`.  These are the supporting files for migration that are too large (>100MB) for the repository.  Building the tar gzip and uploading to the repository as taged release asset has no size limitation.

The tar gzip file is a collection of all the zip files used in the migration.  Typically, a migration file uses a single zip file.  The following command is an example how the tar gzip file could be created:

```bash
> tar -czvf geometry.tar.gz /path/to/*.zip
```

# Generate District Geometry: The Python Script

The Python script `master_data_geom.py` reads geojson from ARC GIS services, defining District boundaries, and generates a SQL file that can be used to INSERT that data into a Postgres database.  There are four URLs defining different missions:

    - Civil Works
    - Regulatory
    - FUDS
    - Military

Not every District has these boundaries defined.  SQL files are generated if that District has geojson defined for a particular mission, and that file is written to a zip file.  The result of running this script is a collection of SQL files, one for each District, each having their respective boundaries defined in the `mission_geom.zip` file.
