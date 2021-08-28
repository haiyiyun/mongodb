package geometry

type GeometryCollection struct {
	Type        string      `json:"type" bson:"type" map:"type"` //default: GeometryCollection
	Coordinates interface{} `json:"coordinates" bson:"coordinates" map:"coordinates"`
}

func NewGeometryCollection(coordinates interface{}) GeometryCollection {
	return GeometryCollection{
		Type:        TypeGeometryCollection,
		Coordinates: coordinates,
	}
}
