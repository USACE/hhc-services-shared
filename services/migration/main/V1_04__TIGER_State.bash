#!/usr/bash

set -e -u

# tiger data is gis_admin and I am hhc schema
PGPASSWORD=${FLYWAY_PASSWORD} psql --host=${PGHOST} --port=5432 --username=${FLYWAY_USER} --dbname=${FP__flyway_database__} \
    --file="/tmp/tiger_state.sql"
