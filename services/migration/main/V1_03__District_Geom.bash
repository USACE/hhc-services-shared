#!/usr/bash

# loop through the files in the mission_geom.zip file
# use those files in this migration

for sqlfile in ${OPT_MIGRATION}/dist_geom/*.sql; do
    echo "executing $sqlfile"

    PGPASSWORD=${FLYWAY_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=${FP__flyway_user__} --dbname=${FP__flyway_database__} \
        --file="$sqlfile"
done
