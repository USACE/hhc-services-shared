-- pgFormatter-ignore
-- ignore the formatter to not format the flyway placeholders

-- Always re-apply roles  when running migrations: ${flyway:timestamp}

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

ALTER ROLE hhc_user SET search_path = ${flyway:defaultSchema}, "tiger", "tiger_data", "public";

GRANT USAGE ON SCHEMA ${flyway:defaultSchema} TO hhc_user;

GRANT hhc_reader , hhc_writer TO hhc_user;

