#!/usr/bash

cd $OPT_MIGRATION

tar -xzvf tl_2024_us_county.tar.gz

shp2pgsql -s 4269 -c -D -w tl_2024_us_county tiger_data.county >/tmp/tiger_county.sql

# tiger data is gis_admin and I am hhc schema
PGPASSWORD=${FLYWAY_PASSWORD} psql --host=${PGHOST} --port=5432 --username=${FLYWAY_USER} --dbname=${FP__flyway_database__} \
    --file="/tmp/tiger_county.sql"
