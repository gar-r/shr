package main

type Url struct {
	Id string `json:"id" bson:"_id"`
	V  string `json:"v" bson:"v"`
}

type Sequence struct {
	S int64 `json:"s" bson:"s"`
}
