package mongodb

import (
	"container/ring"
	"context"
	"sync"

	"go.haiyiyun.org/log"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRingPool struct {
	maxPoolSize int
	mongoRing   *ring.Ring
	mutex       sync.Mutex
	name        string
	opts        []*options.ClientOptions
}

func NewMongoRingPool(name string, maxPoolSize int, opts ...*options.ClientOptions) *MongoRingPool {
	m := &MongoRingPool{
		maxPoolSize: maxPoolSize,
		mongoRing:   ring.New(maxPoolSize),
		name:        name,
		opts:        opts,
	}

	log.Debug("<NewMongoRingPool> init mongo")
	m.createMongos()

	return m
}

func (m *MongoRingPool) createMongos() {
	for i := 0; i < m.maxPoolSize; i++ {
		m.mongoRing.Value = NewMongo(m.name, m.opts...)
		m.mongoRing = m.mongoRing.Next()
	}
}

func (m *MongoRingPool) M() *Mongo {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	mon := m.mongoRing.Value.(*Mongo)
	m.mongoRing = m.mongoRing.Next()

	return mon
}

func (m *MongoRingPool) Disconnect(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for i := 0; i < m.maxPoolSize; i++ {
		m.mongoRing.Value.(*Mongo).Disconnect(ctx)
		m.mongoRing = m.mongoRing.Next()
	}

	log.Debug("<MongoRingPool> disconnect all client of mongodb")
}
