package driver

import (
	"go.haiyiyun.org/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type Driver struct {
	mgo mongodb.Mongoer
	col bson.M
}

func NewDriver(mgo mongodb.Mongoer, col bson.M) *Driver {
	return &Driver{
		mgo: mgo,
		col: col,
	}
}

func (d *Driver) GetMgo() (mgo mongodb.Mongoer) {
	return d.mgo
}
