package driver

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *Driver) ProjectionOne(bd bson.D) *options.FindOneOptions {
	return options.FindOne().SetProjection(bd)
}

func (d *Driver) ProjectionOneID() *options.FindOneOptions {
	return d.ProjectionOne(bson.D{
		{"_id", 1},
	})
}
