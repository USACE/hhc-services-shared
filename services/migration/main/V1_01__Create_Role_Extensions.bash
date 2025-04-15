#!/usr/bash

# all commands are for default schema

# create the extensions
PGPASSWORD=${FLYWAY_PLACEHOLDERS_GIS_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=gis_admin --dbname=${FP__flyway_database__} \
    --command="CREATE EXTENSION IF NOT EXISTS postgis;
        CREATE EXTENSION IF NOT EXISTS postgis_raster;
        CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;
        CREATE EXTENSION IF NOT EXISTS postgis_tiger_geocoder;
        CREATE EXTENSION IF NOT EXISTS postgis_topology;
        CREATE EXTENSION IF NOT EXISTS address_standardizer_data_us;"

# alter the owner
PGPASSWORD=${FLYWAY_PLACEHOLDERS_GIS_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=gis_admin --dbname=${FP__flyway_database__} \
    --command="ALTER SCHEMA tiger OWNER TO gis_admin;
        ALTER SCHEMA tiger_data OWNER TO gis_admin;
        ALTER SCHEMA topology OWNER TO gis_admin;"

# create a function
PGPASSWORD=${FLYWAY_PLACEHOLDERS_GIS_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=gis_admin --dbname=${FP__flyway_database__} \
    --command="CREATE FUNCTION exec(text) returns text language plpgsql volatile AS \$f\$ BEGIN EXECUTE \$1; RETURN \$1; END; \$f\$;"

# run the exec function
PGPASSWORD=${FLYWAY_PLACEHOLDERS_GIS_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=gis_admin --dbname=${FP__flyway_database__} \
    --command="SELECT exec('ALTER TABLE ' || quote_ident(s.nspname) || '.' || quote_ident(s.relname) || ' OWNER TO gis_admin;')
        FROM (
            SELECT nspname, relname
            FROM pg_class c JOIN pg_namespace n ON (c.relnamespace = n.oid) 
            WHERE nspname in ('tiger','topology') AND
            relkind IN ('r','S','v') ORDER BY relkind = 'S')
        s;"

# upgrade
PGPASSWORD=${FLYWAY_PLACEHOLDERS_GIS_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=gis_admin --dbname=${FP__flyway_database__} \
    --command="SELECT postGIS_extensions_upgrade();"
