package model

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/gofrs/uuid"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Office struct {
	ID         string `db:"id" json:"id"`
	Code       string `db:"code" json:"code"`
	Symbol     string `db:"symbol" json:"symbol"`
	Fullname   string `db:"fullname" json:"fullname"`
	OfficeType string `db:"office_type" json:"office_type"`
}

type OfficeAor struct {
	OfficeAorCw   *uuid.UUID `json:"cw"  `
	OfficeAorFuds *uuid.UUID `json:"fuds"`
	OfficeAorMil  *uuid.UUID `json:"mil" `
	OfficeAorReg  *uuid.UUID `json:"reg" `
}

type OfficeFull struct {
	Office
	OfficeAor    *OfficeAor `db:"office_aor" json:"office_aor,omitempty"`
	ParentOffice *Office    `db:"parent_office" json:"parent_office,omitempty"`
}

// ListOffices
func ListOffices(db *pgxpool.Pool, cnames []string) ([]OfficeFull, error) {
	// adding the base column names and any from inputs
	col := []string{"id", "code", "symbol", "fullname", "office_type"}
	col = append(col, cnames...)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(col...)
	sb.From("v_offices_with_geom_id")

	sql, _ := sb.Build()

	var offices []OfficeFull
	if err := pgxscan.Select(context.TODO(), db, &offices, sql); err != nil {
		return nil, err
	}

	return offices, nil
}

// OfficeGeometry
func OfficeGeometry(db *pgxpool.Pool, office string, aors []string) (FeatureCollection, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "code", "symbol", "fullname", "office_type", "srid", "aor", "geom_id", "geometry")
	sb.From("m_office_geojson_3857")
	sb.Where(
		sb.Equal("code", office),
	)

	switch len(aors) {
	case 1:
		sb.Where(sb.Equal("aor", aors[0]))
	case 2:
		sb.Where(sb.Or(sb.Equal("aor", aors[0]), sb.Equal("aor", aors[1])))
	case 3:
		sb.Where(sb.Or(sb.Equal("aor", aors[0]), sb.Equal("aor", aors[1]), sb.Equal("aor", aors[2])))
	case 4:
		sb.Where(sb.Or(sb.Equal("aor", aors[0]), sb.Equal("aor", aors[1]), sb.Equal("aor", aors[2]), sb.Equal("aor", aors[3])))
	}

	sql, args := sb.Build()

	fc := DefaultFeatureCollection()

	rows, err := db.Query(context.TODO(), sql, args...)
	if err != nil {
		return fc, err
	}
	defer rows.Close()

	features := make([]Feature, 0)
	for rows.Next() {
		var feature Feature
		feature.Type = "Feature"
		rows.Scan(&feature.Properties.ID, &feature.Properties.Code,
			&feature.Properties.Symbol,
			&feature.Properties.Fullname,
			&feature.Properties.OfficeType,
			&feature.Properties.SRID,
			&feature.Properties.AOR,
			&feature.Properties.GeomId,
			&feature.Geometry,
		)
		features = append(features, feature)
	}

	fc.Features = features

	return fc, nil

}
