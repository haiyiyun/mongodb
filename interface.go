package mongodb

import (
	"context"
)

type Mongoer interface {
	M() *Mongo
	Disconnect(ctx context.Context)
}
