#!/usr/bash

# loop through the files in the mission_geom.zip file
# use those files in this migration

for sqlfile in ${OPT_MIGRATION}/dist_geom/*.sql; do
    echo "executing $sqlfile"

    PGPASSWORD=${FLYWAY_PASSWORD} psql --host=${PGHOST} --port=5432 --username=${FLYWAY_USER} --dbname=${FP__flyway_database__} \
        --file="$sqlfile"
done
