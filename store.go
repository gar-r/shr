package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connStr = "mongodb://db:27017"
var db = "shr"

type Store interface {
	Find(string)
	Save(Url)
}

type MongoStore struct {
	Session *mongo.Client
	Db *mongo.Database
}

func NewMongoStore(connStr, db string) (*MongoStore, error) {
	opt := options.Client().ApplyURI(connStr)
	client, err := mongo.NewClient(opt); err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err = client.Connect(ctx); err == nil {
			return &MongoStore{
				Session: client,
				Db: client.Database(db),
			}, nil
		}
	}
	return nil, err
}

// func (m *MongoStore) Find(id string) {

// }

// func (m *MongoStore) Save(url Url) {

// }

// func (m *MongoStore) next (int64, error) {
// 	m.Db.Collection("sequence").FindOneAndUpdate()
// }