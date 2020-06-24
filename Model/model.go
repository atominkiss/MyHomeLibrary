package model // nolint:golint

import "labix.org/v2/mgo/bson"

type Book struct {
	ID     bson.ObjectId `bson:"_id"          json:"id"`
	Title  string        `bson:"title"       json:"title"`
	Author *Author       `bson:"author"      json:"author"`
}

type Author struct {
	Firstname string `bson:"firstname"  json:"firstname"`
	Lastname  string `bson:"lastname"   json:"lastname"`
}
