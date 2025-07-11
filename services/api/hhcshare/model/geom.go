package model

import "github.com/gofrs/uuid"

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string       `db:"-" json:"type"`
	ID         string       `db:"id" json:"id"`
	Code       string       `db:"code" json:"code"`
	Symbol     string       `db:"symbol" json:"symbol"`
	Fullname   string       `db:"fullname" json:"fullname"`
	OfficeType string       `db:"office_type" json:"office_type"`
	SRID       int          `db:"srid" json:"srid"`
	AOR        string       `db:"aor" json:"aor"`
	GeomId     uuid.UUID    `db:"geom_id" json:"-"`
	Geometry   MultiPolygon `db:"geometry" json:"geometry"`
}

type Properties struct {
	ID         string    `db:"id" json:"id"`
	Code       string    `db:"code" json:"code"`
	Symbol     string    `db:"symbol" json:"symbol"`
	Fullname   string    `db:"fullname" json:"fullname"`
	OfficeType string    `db:"office_type" json:"office_type"`
	SRID       int       `db:"srid" json:"srid"`
	AOR        string    `db:"-" json:"aor"`
	GeomId     uuid.UUID `db:"geom_id" json:"-"`
}

type Geometry struct {
	Type        string `json:"type"`
	Coordinates any    `json:"coordinates"` // Can be a slice of floats or nested slices
}

type Point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type LineString struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type MultiPolygon struct {
	Type        string          `json:"type"`
	Coordinates [][][][]float64 `json:"coordinates"`
}

// DefaultFeatureCollection
func DefaultFeatureCollection() FeatureCollection {
	var fc FeatureCollection
	fc.Type = "FeatureCollection"
	return fc
}

// DefaultFeature
func DefaultFeature() Feature {
	var f Feature
	f.Type = "Feature"
	return f
}
