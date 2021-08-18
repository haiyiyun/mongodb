package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoPool(poolType, name string, maxPoolSize int, opts ...*options.ClientOptions) (mr Mongoer) {
	switch poolType {
	case "ring":
		mr = NewMongoRingPool(name, maxPoolSize, opts...)
	case "chan":
		mr = NewMongoChanPool(name, maxPoolSize, opts...)
	default:
		mr = NewMongo(name, opts...)
	}

	return
}
