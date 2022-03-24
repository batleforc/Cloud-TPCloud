package model

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DbPrefixName = "TpCloud"
)

type Db struct {
	DbName string
	dbUri  string
	option *options.ClientOptions
}

func (v *Db) Init(uri string, Env string) *Db {
	v.DbName = DbPrefixName + Env
	v.dbUri = uri
	v.option = options.Client().ApplyURI(v.dbUri)
	return v
}
func (v *Db) GetClient() (*mongo.Client, error) {
	return mongo.NewClient(v.option)
}

func (v *Db) GetDb(client *mongo.Client) *mongo.Database {
	return client.Database(v.DbName)
}
