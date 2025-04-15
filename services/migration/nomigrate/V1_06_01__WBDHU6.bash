#!/usr/bash

# create the table before loading data
# PGPASSWORD=${FLYWAY_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=${FP__flyway_user__} --dbname=${FP__flyway_database__} \
#     --commands="CREATE TABLE usace.wbdhu6 (gid serial, objectid numeric, areaacres numeric, states varchar(50), huc6 varchar(6), name varchar(120), shape_leng numeric, shape_area numeric;
#         ALTER TABLE usace.wbdhu6 ADD PRIMARY KEY (gid);
#         SELECT AddGeometryColumn('usace' , 'wbdhu6' , 'geom' , '4269' , 'MULTIPOLYGON' , 2);
#         COMMIT;
#         ANALYZE usace.wbdhu6;"


# echo "CREATE TABLE usace.wbdhu6 (gid serial, objectid numeric, areaacres numeric, states varchar(50), huc6 varchar(6), name varchar(120), shape_leng numeric, shape_area numeric);" >/tmp/wbdhu6.sql
# echo "ALTER TABLE usace.wbdhu6 ADD PRIMARY KEY (gid);" >>/tmp/wbdhu6.sql
# echo "COMMIT;" >>/tmp/wbdhu6.sql
# echo "ANALYZE usace.wbdhu6;" >>/tmp/wbdhu6.sql
# echo "SELECT AddGeometryColumn('usace' , 'wbdhu6' , 'geom' , '4269' , 'MULTIPOLYGON' , 2);" >>/tmp/wbdhu6.sql

cd $OPT_MIGRATION
unzip -o wbdhu6.zip

shp2pgsql -s 4269 -c -D -w wbdhu6 usace.wbdhu6 >/tmp/wbdhu6.sql

# tiger data is gis_admin and I am usace schema
PGPASSWORD=${FLYWAY_PASSWORD} psql --host=${FLYWAY_DB_HOST} --port=5432 --username=${FP__flyway_user__} --dbname=${FP__flyway_database__} \
    --file="/tmp/wbdhu6.sql"
