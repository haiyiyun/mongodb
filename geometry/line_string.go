package geometry

type LineString struct {
	Type        string             `json:"type" bson:"type" map:"type"` //default: LineString
	Coordinates []PointCoordinates `json:"coordinates" bson:"coordinates" map:"coordinates"`
}

func NewLineString(coordinates []PointCoordinates) LineString {
	return LineString{
		Type:        TypeLineString,
		Coordinates: coordinates,
	}
}
