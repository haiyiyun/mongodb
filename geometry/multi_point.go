package geometry

type MultiPoint struct {
	Type        string             `json:"type" bson:"type" map:"type"` //default: MultiPoint
	Coordinates []PointCoordinates `json:"coordinates" bson:"coordinates" map:"coordinates"`
}

func NewMultiPoint(coordinates []PointCoordinates) MultiPoint {
	return MultiPoint{
		Type:        TypeMultiPoint,
		Coordinates: coordinates,
	}
}
