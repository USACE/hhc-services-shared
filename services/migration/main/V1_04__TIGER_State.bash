#!/usr/bash

cd $OPT_MIGRATION

unzip -o tl_2024_us_state.zip

shp2pgsql -s 4269 -c -D -w tl_2024_us_state tiger_data.state >/tmp/tiger_state.sql

# tiger data is gis_admin and I am usace schema
PGPASSWORD=${FLYWAY_PLACEHOLDERS_GIS_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=gis_admin --dbname=${FP__flyway_database__} \
    --file="/tmp/tiger_state.sql"
