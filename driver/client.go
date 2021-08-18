package driver

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//如果在DNS处定义了相关concern的options，则无需StartTransaction时设置，否则尽量在StartTransaction时设置相关concern的options
//例如：
//	err := sctx.StartTransaction(options.Transaction().SetReadConcern(readconcern.Majority()).SetWriteConcern(writeconcern.New(writeconcern.WMajority())))
func (d *Driver) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return d.mgo.M().Client().UseSession(ctx, fn)
}
