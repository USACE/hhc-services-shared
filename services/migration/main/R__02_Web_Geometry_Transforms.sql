-- pg Formatter-ignore
-- ignore the formatter to not format the flyway placeholders
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- materialized views providing transformed geometries
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- hhc.office_aor_cw
CREATE MATERIALIZED VIEW IF NOT EXISTS m_office_aor_cw_3857 AS
SELECT
    office_id
    , public.ST_Transform(geom , 3857) AS geom
FROM
    hhc.office_aor_cw;

-- hhc.office_aor_reg
CREATE MATERIALIZED VIEW IF NOT EXISTS m_office_aor_reg_3857 AS
SELECT
    office_id
    , public.ST_Transform(geom , 3857) AS geom
FROM
    hhc.office_aor_reg;

-- hhc.office_aor_fuds
CREATE MATERIALIZED VIEW IF NOT EXISTS m_office_aor_fuds_3857 AS
SELECT
    office_id
    , public.ST_Transform(geom , 3857) AS geom
FROM
    hhc.office_aor_fuds;

-- hhc.office_aor_mil
CREATE MATERIALIZED VIEW IF NOT EXISTS m_office_aor_mil_3857 AS
SELECT
    office_id
    , public.ST_Transform(geom , 3857) AS geom
FROM
    hhc.office_aor_mil;

