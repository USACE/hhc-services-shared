-- pgFormatter-ignore
-- ignore the formatter to not format the flyway placeholders

-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- materialized views providing transformed geometries
-- *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*
-- hhc.v_office_with_geom
CREATE OR REPLACE VIEW v_offices_with_geom_id AS
SELECT
	o.id
	, o.code
	, o.symbol
	, o.fullname
	, o.office_type
	, json_build_object('cw' , cw.id , 'fuds' , fuds.id , 'mil' , mil.id ,
	'reg' , reg.id) AS office_aor
	, json_build_object('id' , p.id , 'code' , p.code , 'symbol' , p.symbol ,
	'fullname' , p.fullname , 'office_type' , p.office_type) AS parent_office
FROM
	hhc.office o
	LEFT JOIN hhc.office p ON p.id = o.parent_id
	LEFT JOIN hhc.office_aor_cw cw ON cw.office_id = o.id
	LEFT JOIN hhc.office_aor_fuds fuds ON fuds.office_id = o.id
	LEFT JOIN hhc.office_aor_mil mil ON mil.office_id = o.id
	LEFT JOIN hhc.office_aor_reg reg ON reg.office_id = o.id
WHERE
	o.parent_id IS NOT NULL
	AND o.office_type IN ('DIST' , 'MSC')
ORDER BY
	o.id;

GRANT SELECT ON ${flyway:defaultSchema}.v_offices_with_geom_id TO hhc_shared_reader;
