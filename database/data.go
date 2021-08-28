package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *Database) DataSet(data bson.D) bson.D {
	return bson.D{
		{"$set", data},
	}
}

func (db *Database) DataUnSet(data bson.D) bson.D {
	return bson.D{
		{"$unset", data},
	}
}

func (db *Database) DataAddToSet(data bson.D) bson.D {
	return bson.D{
		{"$addToSet", data},
	}
}

func (db *Database) DataPull(data bson.D) bson.D {
	return bson.D{
		{"$pull", data},
	}
}

func (db *Database) DataPullAll(data bson.D) bson.D {
	return bson.D{
		{"$pullAll", data},
	}
}

func (db *Database) DataPush(data bson.D) bson.D {
	return bson.D{
		{"$push", data},
	}
}

func (db *Database) DataPop(data bson.D) bson.D {
	return bson.D{
		{"$pop", data},
	}
}

func (db *Database) DataEach(data bson.D) bson.D {
	return bson.D{
		{"$each", data},
	}
}

func (db *Database) DataInc(data bson.D) bson.D {
	return bson.D{
		{"$inc", data},
	}
}

func (db *Database) DataSetOnInsert(data bson.D, opts ...*options.UpdateOptions) (bson.D, *options.UpdateOptions) {
	return bson.D{
		bson.E{"$setOnInsert", data},
	}, options.MergeUpdateOptions(opts...).SetUpsert(true)
}
