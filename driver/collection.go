package driver

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *Driver) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return d.mgo.M().Col(d.col).BulkWrite(ctx, models, opts...)
}

func (d *Driver) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return d.mgo.M().Col(d.col).InsertMany(ctx, documents, opts...)
}

func (d *Driver) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return d.mgo.M().Col(d.col).InsertOne(ctx, document, opts...)
}

func (d *Driver) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return d.mgo.M().Col(d.col).DeleteOne(ctx, filter, opts...)
}

func (d *Driver) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return d.mgo.M().Col(d.col).DeleteMany(ctx, filter, opts...)
}

func (d *Driver) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return d.mgo.M().Col(d.col).Find(ctx, filter, opts...)
}

func (d *Driver) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return d.mgo.M().Col(d.col).FindOne(ctx, filter, opts...)
}

func (d *Driver) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	return d.mgo.M().Col(d.col).FindOneAndDelete(ctx, filter, opts...)
}

func (d *Driver) FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	return d.mgo.M().Col(d.col).FindOneAndReplace(ctx, filter, replacement, opts...)
}

func (d *Driver) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return d.mgo.M().Col(d.col).FindOneAndUpdate(ctx, filter, update, opts...)
}

func (d *Driver) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return d.mgo.M().Col(d.col).ReplaceOne(ctx, filter, replacement, opts...)
}

func (d *Driver) UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return d.mgo.M().Col(d.col).UpdateByID(ctx, id, update, opts...)
}

func (d *Driver) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return d.mgo.M().Col(d.col).UpdateMany(ctx, filter, update, opts...)
}

func (d *Driver) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return d.mgo.M().Col(d.col).UpdateOne(ctx, filter, update, opts...)
}

func (d *Driver) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return d.mgo.M().Col(d.col).Distinct(ctx, fieldName, filter, opts...)
}

func (d *Driver) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return d.mgo.M().Col(d.col).Aggregate(ctx, pipeline, opts...)
}

func (d *Driver) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return d.mgo.M().Col(d.col).CountDocuments(ctx, filter, opts...)
}

func (d *Driver) EstimatedDocumentCount(ctx context.Context, filter bson.D, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return d.mgo.M().Col(d.col).EstimatedDocumentCount(ctx, opts...)
}
