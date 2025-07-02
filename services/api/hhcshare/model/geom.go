package model

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string       `db:"-" json:"type"`
	Properties Properties   `db:"properties" json:"properties"`
	Geometry   MultiPolygon `db:"geometry" json:"geometry"`
}

type Properties struct {
	Office
	SRID int `db:"srid" json:"srid"`
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
