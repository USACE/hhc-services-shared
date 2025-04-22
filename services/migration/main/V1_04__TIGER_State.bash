#!/usr/bash

cd $OPT_MIGRATION

unzip -o tl_2024_us_state.zip

shp2pgsql -s 4269 -c -D -w tl_2024_us_state tiger_data.state >/tmp/tiger_state.sql

# tiger data is gis_admin and I am hhc schema
PGPASSWORD=${FLYWAY_PASSWORD} psql --host=${PGHOST} --port=5432 --username=${FLYWAY_USER} --dbname=${FP__flyway_database__} \
    --file="/tmp/tiger_state.sql"
