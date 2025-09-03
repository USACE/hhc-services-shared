-- pgFormatter-ignore

-- remove the materialized view m_office_geojson_3857 and
-- view v_offices_with_geom_id before altering

-- these views will be added back in the repeatable migrations

DROP MATERIALIZED VIEW IF EXISTS m_office_geojson_3857;

DROP VIEW IF EXISTS v_offices_with_geom_id;

-- remove the character length restrictions on the office table
-- for columns 'code' and 'symbol'
ALTER TABLE office
    ALTER COLUMN code TYPE VARCHAR,
    ALTER COLUMN symbol TYPE VARCHAR;

-- populate the office table with Other Offices
INSERT INTO office(id , code , symbol , fullname , office_type , parent_id)
    VALUES ('Q0' , 'IWR' , 'CEIWR' , 'Institute For Water Resources' , 'OTHER' , 'S0')
    ,('U4' , 'ERDC' , 'CEERDC' , 'Engineer Research and Development Center' , 'OTHER' , 'S0');
