-- pgFormatter-ignore
-- ignore the formatter to not format the flyway placeholders

-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- materialized views providing transformed geometries
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- hhc.office_aor_cw
CREATE MATERIALIZED VIEW IF NOT EXISTS m_office_aor_cw_3857 AS
SELECT
    id
    , office_id
    , public.ST_Transform(geom , 3857) AS geom
FROM
    hhc.office_aor_cw;

GRANT SELECT ON ${flyway:defaultSchema}.m_office_aor_cw_3857 TO hhc_shared_reader;

-- hhc.office_aor_reg
CREATE MATERIALIZED VIEW IF NOT EXISTS m_office_aor_reg_3857 AS
SELECT
    id
    , office_id
    , public.ST_Transform(geom , 3857) AS geom
FROM
    hhc.office_aor_reg;

GRANT SELECT ON ${flyway:defaultSchema}.m_office_aor_reg_3857 TO hhc_shared_reader;

-- hhc.office_aor_fuds
CREATE MATERIALIZED VIEW IF NOT EXISTS m_office_aor_fuds_3857 AS
SELECT
    id
    , office_id
    , public.ST_Transform(geom , 3857) AS geom
FROM
    hhc.office_aor_fuds;

GRANT SELECT ON ${flyway:defaultSchema}.m_office_aor_fuds_3857 TO hhc_shared_reader;

-- hhc.office_aor_mil
CREATE MATERIALIZED VIEW IF NOT EXISTS m_office_aor_mil_3857 AS
SELECT
    id
    , office_id
    , public.ST_Transform(geom , 3857) AS geom
FROM
    hhc.office_aor_mil;

GRANT SELECT ON ${flyway:defaultSchema}.m_office_aor_mil_3857 TO hhc_shared_reader;

-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- materialized views providing transformed geometries
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- office geojson with web mercator
CREATE MATERIALIZED VIEW IF NOT EXISTS m_office_geojson_3857 AS
SELECT
    o.id
    , o.code
    , o.symbol
    , o.fullname
    , o.office_type
    , 3857 AS srid
    , 'civil_works' AS aor
    , oac.id AS geom_id
    , ST_AsGeoJSON(ST_Transform(oac.geom , 3857))::json AS geometry
FROM
    office o
    JOIN office_aor_cw oac ON oac.office_id = o.id
UNION ALL
SELECT
    o.id
    , o.code
    , o.symbol
    , o.fullname
    , o.office_type
    , 3857 AS srid
    , 'fuds' AS aor
    , oac.id AS geom_id
    , ST_AsGeoJSON(ST_Transform(oac.geom , 3857))::json AS geometry
FROM
    office o
    JOIN office_aor_fuds oac ON oac.office_id = o.id
UNION ALL
SELECT
    o.id
    , o.code
    , o.symbol
    , o.fullname
    , o.office_type
    , 3857 AS srid
    , 'military' AS aor
    , oac.id AS geom_id
    , ST_AsGeoJSON(ST_Transform(oac.geom , 3857))::json AS geometry
FROM
    office o
    JOIN office_aor_mil oac ON oac.office_id = o.id
UNION ALL
SELECT
    o.id
    , o.code
    , o.symbol
    , o.fullname
    , o.office_type
    , 3857 AS srid
    , 'regulatory' AS aor
    , oac.id AS geom_id
    , ST_AsGeoJSON(ST_Transform(oac.geom , 3857))::json AS geometry
FROM
    office o
    JOIN office_aor_reg oac ON oac.office_id = o.id;

GRANT SELECT ON ${flyway:defaultSchema}.m_office_geojson_3857 TO hhc_shared_reader;

