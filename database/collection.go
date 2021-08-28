package database

import (
	"context"
	"time"

	"github.com/haiyiyun/utils/help"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *Database) UpdateOne(ctx context.Context, filter, update bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	update = append(update, db.DataSet(bson.D{
		{"update_time", time.Now()},
	})...)

	return db.Driver.UpdateOne(ctx, filter, update, opts...)
}

func (db *Database) Set(ctx context.Context, filter, data bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.UpdateOne(ctx, filter, db.DataSet(data), opts...)
}

func (db *Database) UnSet(ctx context.Context, filter, data bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.UpdateOne(ctx, filter, db.DataUnSet(data), opts...)
}

func (db *Database) SetOnInsert(ctx context.Context, filter, update bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	insertData, opt := db.DataSetOnInsert(bson.D{
		{"create_time", time.Now()},
	}, opts...)

	return db.UpdateOne(ctx, filter, append(update, insertData...), opt)
}

func (db *Database) SetAndSetOnInsert(ctx context.Context, filter, setData bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	insertData, opt := db.DataSetOnInsert(bson.D{
		{"create_time", time.Now()},
	}, opts...)

	return db.UpdateOne(ctx, filter, append(db.DataSet(setData), insertData...), opt)
}

func (db *Database) AddToSet(ctx context.Context, filter, addToSetData bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.UpdateOne(ctx, filter, db.DataAddToSet(addToSetData), opts...)
}

func (db *Database) AddToSetOnInsert(ctx context.Context, filter, addToSetData bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.SetOnInsert(ctx, filter, db.DataAddToSet(addToSetData), opts...)
}

func (db *Database) SetAndAddToSet(ctx context.Context, filter, setData, addToSetData bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.UpdateOne(ctx, filter, append(db.DataSet(setData), db.DataAddToSet(addToSetData)...), opts...)
}

func (db *Database) SetAndAddToSetOnInsert(ctx context.Context, filter, setData, addToSetData bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.SetOnInsert(ctx, filter, append(db.DataSet(setData), db.DataAddToSet(addToSetData)...), opts...)
}

func (db *Database) Pull(ctx context.Context, filter, pullFilters bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.UpdateOne(ctx, filter, db.DataPull(pullFilters), opts...)
}

func (db *Database) Push(ctx context.Context, filter, pushFilters bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.UpdateOne(ctx, filter, db.DataPush(pushFilters), opts...)
}

func (db *Database) Create(ctx context.Context, m interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	mm := help.NewStruct(m).StructToMap()

	if id, found := mm["_id"]; found {
		if id == primitive.NilObjectID {
			mm["_id"] = primitive.NewObjectID()
		}
	}

	if ct, found := mm["create_time"].(time.Time); found {
		if ct.IsZero() {
			mm["create_time"] = time.Now()
		}
	}

	if ut, found := mm["update_time"].(time.Time); found {
		if ut.IsZero() {
			mm["update_time"] = time.Now()
		}
	}

	return db.InsertOne(ctx, mm, opts...)
}
