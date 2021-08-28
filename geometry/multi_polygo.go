package geometry

type MultiPolygon struct {
	Type        string                 `json:"type" bson:"type" map:"type"` //default: MultiPolygon
	Coordinates [][][]PointCoordinates `json:"coordinates" bson:"coordinates" map:"coordinates"`
}

func NewMultiPolygon(coordinates [][][]PointCoordinates) MultiPolygon {
	return MultiPolygon{
		Type:        TypeMultiPolygon,
		Coordinates: coordinates,
	}
}
