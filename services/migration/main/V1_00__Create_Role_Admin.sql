-- pgFormatter-ignore
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- create the gis_admin role
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
DO $$
BEGIN
    CREATE USER gis_admin WITH ENCRYPTED PASSWORD '${GIS_PASSWORD}';
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating role gis_admin -- it already exists';
END
$$;

ALTER ROLE gis_admin WITH SUPERUSER;

ALTER ROLE gis_admin SET search_path = ${flyway:defaultSchema}, "public" , "tiger", "tiger_data";

GRANT ALL PRIVILEGES ON DATABASE postgres TO gis_admin;

-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- create the hhc_user, hhc_reader, and hhc_writer roles
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
DO $$
BEGIN
    CREATE USER hhc_user WITH ENCRYPTED PASSWORD '${APP_PASSWORD}';
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating role hhc_user -- it already exists';
END
$$;

DO $$
BEGIN
    CREATE ROLE hhc_reader;
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating role hhc_reader -- it already exists';
END
$$;

-- Role hhc_writer
DO $$
BEGIN
    CREATE ROLE hhc_writer;
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating role hhc_writer -- it already exists';
END
$$;

-- GRANT for default schema roles
GRANT SELECT ON ALL TABLES IN SCHEMA ${flyway:defaultSchema} TO hhc_reader;

GRANT INSERT , UPDATE , DELETE ON ALL TABLES IN SCHEMA ${flyway:defaultSchema} TO hhc_writer;

REVOKE ALL ON flyway_schema_history FROM hhc_reader;

REVOKE ALL ON flyway_schema_history FROM hhc_writer;

ALTER ROLE hhc_user SET search_path = ${flyway:defaultSchema}, "public" , "tiger", "tiger_data";

GRANT USAGE ON SCHEMA ${flyway:defaultSchema} TO hhc_user;

GRANT hhc_reader , hhc_writer TO hhc_user;

