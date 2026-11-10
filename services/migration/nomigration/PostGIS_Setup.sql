-- PostGIS Extension Setup not happening in newer PostGIS images

-- Role
DO $$
BEGIN
  IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'gis_admin') THEN
    CREATE ROLE gis_admin;
  END IF;
END
$$;

-- Extensions (these create the tiger and topology schemas)
CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;
CREATE EXTENSION IF NOT EXISTS postgis_tiger_geocoder;
CREATE EXTENSION IF NOT EXISTS postgis_topology;

-- Schemas and ownership
CREATE SCHEMA IF NOT EXISTS tiger_data AUTHORIZATION gis_admin;
ALTER SCHEMA tiger      OWNER TO gis_admin;
ALTER SCHEMA tiger_data OWNER TO gis_admin;
ALTER SCHEMA topology   OWNER TO gis_admin;
