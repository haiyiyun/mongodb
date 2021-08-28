package geometry

type MultiLineString struct {
	Type        string               `json:"type" bson:"type" map:"type"` //default: MultiLineString
	Coordinates [][]PointCoordinates `json:"coordinates" bson:"coordinates" map:"coordinates"`
}

func NewMultiLineString(coordinates [][]PointCoordinates) MultiLineString {
	return MultiLineString{
		Type:        TypeMultiLineString,
		Coordinates: coordinates,
	}
}
