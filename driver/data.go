package driver

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *Driver) DataSet(data bson.D) bson.D {
	return bson.D{
		{"$set", data},
	}
}

func (d *Driver) DataUnSet(data bson.D) bson.D {
	return bson.D{
		{"$unset", data},
	}
}

func (d *Driver) DataAddToSet(data bson.D) bson.D {
	return bson.D{
		{"$addToSet", data},
	}
}

func (d *Driver) DataPull(data bson.D) bson.D {
	return bson.D{
		{"$pull", data},
	}
}

func (d *Driver) DataInc(data bson.D) bson.D {
	return bson.D{
		{"$inc", data},
	}
}

func (d *Driver) DataSetOnInsert(data bson.D, opts ...*options.UpdateOptions) (bson.D, *options.UpdateOptions) {
	return bson.D{
		bson.E{"$setOnInsert", data},
	}, options.MergeUpdateOptions(opts...).SetUpsert(true)
}
