package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Session *mongo.Client
	Db      *mongo.Database
}

func NewMongoStore(connStr, db string) (store *Store, err error) {
	opt := options.Client().ApplyURI(connStr)
	if client, err := mongo.NewClient(opt); err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err = client.Connect(ctx); err == nil {
			store = &Store{
				Session: client,
				Db:      client.Database(db),
			}
		}
	}
	return
}

func (s *Store) Find(id string) (url *Url, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res := s.Db.Collection("urls").FindOne(ctx, bson.M{"_id": id})
	err = res.Err()
	if err == nil {
		err = res.Decode(&url)
	}
	return
}

func (s *Store) FindByUrl(u string) (url *Url, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res := s.Db.Collection("urls").FindOne(ctx, bson.M{"val": u})
	err = res.Err()
	if err == nil {
		err = res.Decode(&url)
	}
	return
}

func (s *Store) Save(url *Url) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = s.Db.Collection("urls").InsertOne(ctx, url)
	return
}

func (s *Store) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := s.Session.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
