package geometry

type Polygon struct {
	Type        string               `json:"type" bson:"type" map:"type"` //default: Polygon
	Coordinates [][]PointCoordinates `json:"coordinates" bson:"coordinates" map:"coordinates"`
}

func NewPolygon(coordinates [][]PointCoordinates) Polygon {
	return Polygon{
		Type:        TypePolygon,
		Coordinates: coordinates,
	}
}
