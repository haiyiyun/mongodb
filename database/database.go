package database

import (
	"github.com/haiyiyun/mongodb"
	"github.com/haiyiyun/mongodb/driver"
	"go.mongodb.org/mongo-driver/bson"
)

type Database struct {
	*driver.Driver `json:"-" bson:"-" map:"-"`
}

func NewDatabase(mgo mongodb.Mongoer, col bson.M) *Database {
	mdl := &Database{
		Driver: driver.NewDriver(mgo, col),
	}

	return mdl
}
