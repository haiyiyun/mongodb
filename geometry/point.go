package geometry

type Point struct {
	Type        string           `json:"type" bson:"type" map:"type"` //default: Point
	Coordinates PointCoordinates `json:"coordinates" bson:"coordinates" map:"coordinates"`
}

func NewPoint(coordinates PointCoordinates) Point {
	return Point{
		Type:        TypePoint,
		Coordinates: coordinates,
	}
}
