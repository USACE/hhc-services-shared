-- pgFormatter-ignore
-- ignore the formatter to not format the flyway placeholders

-- Always re-apply roles  when running migrations: ${flyway:timestamp}

-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- create the user, reader, and writer roles
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
DO $$
BEGIN
    CREATE USER ${APP_USER} WITH ENCRYPTED PASSWORD '${APP_PASSWORD}';
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating ${APP_USER} role -- it already exists';
END
$$;

DO $$
BEGIN
    CREATE ROLE hhc_shared_reader;
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating reader role -- it already exists';
END
$$;

-- Role hhc_shared_writer
DO $$
BEGIN
    CREATE ROLE hhc_shared_writer;
EXCEPTION
    WHEN DUPLICATE_OBJECT THEN
        RAISE NOTICE 'not creating writer role -- it already exists';
END
$$;

-- GRANT for default schema roles
GRANT SELECT ON ALL TABLES IN SCHEMA ${flyway:defaultSchema} TO hhc_shared_reader;
GRANT INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA ${flyway:defaultSchema} TO hhc_shared_writer;

REVOKE ALL ON flyway_schema_history FROM hhc_shared_reader;
REVOKE ALL ON flyway_schema_history FROM hhc_shared_writer;

GRANT hhc_shared_reader, hhc_shared_writer TO ${APP_USER};

GRANT USAGE ON SCHEMA ${flyway:defaultSchema} TO ${APP_USER};

ALTER ROLE ${APP_USER} SET search_path = public, ${flyway:defaultSchema};

-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- granting for tiger_data with a check the schema exists
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = 'tiger_data') THEN
        EXECUTE 'GRANT SELECT ON ALL TABLES IN SCHEMA tiger_data TO PUBLIC;';
    ELSE
        RAISE EXCEPTION 'Schema hhc does not exist. Grant SELECT TO PUBLIC failed.';
    END IF;
END
$$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = 'tiger_data') THEN
        EXECUTE 'GRANT USAGE ON SCHEMA tiger_data TO PUBLIC;';
    ELSE
        RAISE EXCEPTION 'Schema hhc does not exist. GRANT USAGE failed.';
    END IF;
END
$$;
