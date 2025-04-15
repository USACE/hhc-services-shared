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
-- create the usace_user, usace_reader, and usace_writer roles
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
DO $$
BEGIN
    CREATE USER usace_user WITH ENCRYPTED PASSWORD '${APP_PASSWORD}';
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating role usace_user -- it already exists';
END
$$;

DO $$
BEGIN
    CREATE ROLE usace_reader;
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating role usace_reader -- it already exists';
END
$$;

-- Role usace_writer
DO $$
BEGIN
    CREATE ROLE usace_writer;
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating role usace_writer -- it already exists';
END
$$;

-- GRANT for default schema roles
GRANT SELECT ON ALL TABLES IN SCHEMA ${flyway:defaultSchema} TO usace_reader;

GRANT INSERT , UPDATE , DELETE ON ALL TABLES IN SCHEMA ${flyway:defaultSchema} TO usace_writer;

REVOKE ALL ON flyway_schema_history FROM usace_reader;

REVOKE ALL ON flyway_schema_history FROM usace_writer;

ALTER ROLE usace_user SET search_path = ${flyway:defaultSchema}, "public" , "tiger", "tiger_data";

GRANT USAGE ON SCHEMA ${flyway:defaultSchema} TO usace_user;

GRANT usace_reader , usace_writer TO usace_user;

