package model

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Office struct {
	ID         string `db:"id" json:"id"`
	Code       string `db:"code" json:"code"`
	Symbol     string `db:"symbol" json:"symbol"`
	Fullname   string `db:"fullname" json:"fullname"`
	OfficeType string `db:"office_type" json:"office_type"`
}

type OfficeFull struct {
	Office
	ParentOffice Office `db:"parent_office" json:"parent_office"`
}

// ListOffices
func ListOffices(db *pgxpool.Pool) ([]Office, error) {
	var sql string = `
SELECT
	o.id
	, o.code
	, o.symbol
	, o.fullname
	, o.office_type
FROM
	office o
WHERE
	o.parent_id IS NOT NULL
	AND (o.office_type = 'DIST'
		OR o.office_type = 'MSC')
ORDER BY
	o.id
`
	oo := make([]Office, 0)
	err := pgxscan.Select(context.TODO(), db, &oo, sql)

	return oo, err
}

// ListOfficesFull
func ListOfficesFull(db *pgxpool.Pool) ([]OfficeFull, error) {
	var sql string = `
SELECT
	o.id
	, o.code
	, o.symbol
	, o.fullname
	, o.office_type
	, json_build_object('id' , p.id , 'code' , p.code , 'symbol' , p.symbol ,
	'fullname' , p.fullname, 'office_type', p.office_type ) AS parent_office
FROM
	office o
	JOIN office p ON p.id = o.parent_id
WHERE
	o.parent_id IS NOT NULL
	AND (o.office_type = 'DIST'
		OR o.office_type = 'MSC')
ORDER BY
	o.id
`
	oo := make([]OfficeFull, 0)
	err := pgxscan.Select(context.TODO(), db, &oo, sql)

	return oo, err
}

// OfficeGeometry
func OfficeGeometry(db *pgxpool.Pool, office string) (FeatureCollection, error) {
	var sql string = `
SELECT
	json_build_object('id' , o.id , 'code' , o.code , 'symbol' , o.symbol , 'fullname'
	, o.fullname , 'office_type' , o.office_type , 'srid' , ST_SRID (oac.geom)) AS properties
	, ST_AsGeoJSON (oac.geom)::json AS geometry
FROM
	office o
	JOIN office_aor_cw oac ON oac.office_id = o.id
WHERE
	o.code = $1
`
	fc := DefaultFeatureCollection()

	feature := DefaultFeature()

	err := pgxscan.Get(context.TODO(), db, &feature, sql, office)

	fc.Features = append(fc.Features, feature)

	return fc, err

}
