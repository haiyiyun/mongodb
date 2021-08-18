package mongodb

import (
	"context"

	"go.haiyiyun.org/log"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoChanPool struct {
	maxPoolSize int
	mongoChan   chan *Mongo
	name        string
	opts        []*options.ClientOptions
}

func NewMongoChanPool(name string, maxPoolSize int, opts ...*options.ClientOptions) *MongoChanPool {
	m := &MongoChanPool{
		maxPoolSize: maxPoolSize,
		mongoChan:   make(chan *Mongo, maxPoolSize),
		name:        name,
		opts:        opts,
	}

	log.Debug("<NewMongoChanPool> init mongo")
	m.createMongos()

	return m
}

func (m *MongoChanPool) createMongos() {
	for i := 0; i < m.maxPoolSize; i++ {
		m.mongoChan <- NewMongo(m.name, m.opts...)
	}
}

func (m *MongoChanPool) M() *Mongo {
	mon := <-m.mongoChan
	m.mongoChan <- NewMongo(m.name, m.opts...)
	return mon
}

func (m *MongoChanPool) Disconnect(ctx context.Context) {
	for i := 0; i < len(m.mongoChan); i++ {
		mon := <-m.mongoChan
		mon.Disconnect(ctx)
	}

	log.Debug("<MongoChanPool> disconnect all client of mongodb")
}
