package driver

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *Driver) FilterByID(id primitive.ObjectID) bson.D {
	return bson.D{{"_id", id}}
}
