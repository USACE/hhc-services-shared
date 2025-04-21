#!/usr/bash

cd $OPT_MIGRATION

unzip -o wbdhu6.zip

shp2pgsql -s 4269 -c -D -w wbdhu6 hhc.wbdhu6 >/tmp/wbdhu6.sql

# tiger data is gis_admin and I am hhc schema
PGPASSWORD=${FLYWAY_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=${FP__flyway_user__} --dbname=${FP__flyway_database__} \
    --file="/tmp/wbdhu6.sql"
