package main

type Url struct {
	Id  string `json:"id" bson:"_id"`
	Val string `json:"val" bson:"val"`
}
