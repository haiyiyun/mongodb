package mongodb

import (
	"context"

	"github.com/haiyiyun/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	database    string
	client      *mongo.Client
	collections map[string]*mongo.Collection
	opts        []*options.ClientOptions
}

func NewMongo(name string, opts ...*options.ClientOptions) *Mongo {
	m := &Mongo{
		database:    name,
		collections: map[string]*mongo.Collection{},
		opts:        opts,
	}

	mClient, err := mongo.NewClient(opts...)
	if err != nil {
		log.Fatal("<NewMongo> ", "mongo.NewClient error:", err)
	}

	m.client = mClient
	m.ConnectDefault(context.Background())

	return m
}

func (m *Mongo) M() *Mongo {
	return m
}

func (m *Mongo) Client() *mongo.Client {
	return m.client
}

func (m *Mongo) Col(col bson.M) *mongo.Collection {
	if colName, ok := col["name"].(string); ok {
		if _, found := m.collections[colName]; found {
			return m.collections[colName]
		}
	}

	log.Fatal("First you must use schema of collection to InitCollection then can call Col!")
	return nil
}

func (m *Mongo) InitCollection(col bson.M) {
	if colName, ok := col["name"].(string); ok {
		if _, found := m.collections[colName]; !found {
			mdb := m.client.Database(m.database)
			opts := options.CreateCollection()
			if _, found := col["schema"]; found {
				validator := bson.M{
					"$jsonSchema": col["schema"],
				}
				opts.SetValidator(validator)

				if validationLevel, found := col["validationLevel"]; found && validationLevel.(string) != "" {
					opts.SetValidationLevel(validationLevel.(string))

				}

				if validationAction, found := col["validationAction"]; found && validationAction.(string) != "" {
					opts.SetValidationAction(validationAction.(string))

				}
			}

			if err := mdb.CreateCollection(context.Background(), colName, opts); err != nil {
				if ce, ok := err.(mongo.CommandError); ok {
					if !ce.HasErrorMessage("Collection already exists") {
						log.Fatal("CreateCollection error:", err)
					}
				}
			}

			coll := mdb.Collection(colName)
			if _, found := col["index"]; found {
				if colIndexs, okType := col["index"].([]mongo.IndexModel); okType {
					coll.Indexes().CreateMany(context.Background(), colIndexs)
				}
			}

			m.collections[colName] = coll
		}
	}
}

func (m *Mongo) Connect(ctx context.Context) {
	if err := m.client.Connect(context.Background()); err != nil {
		log.Fatal("Connect error:", err)
	}
}

func (m *Mongo) Ping(ctx context.Context, rp *readpref.ReadPref) {
	if err := m.client.Ping(ctx, rp); err != nil {
		log.Fatal("Ping error:", err)
	}
}

func (m *Mongo) Disconnect(ctx context.Context) {
	m.client.Disconnect(ctx)
}

func (m *Mongo) ConnectDefault(ctx context.Context) {
	m.Connect(ctx)

	//判断服务是否可用
	m.Ping(ctx, readpref.Primary())
}
