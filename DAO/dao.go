package dao

import "labix.org/v2/mgo"

type BooksDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const COLLECTION = "books"

// Establish a connection to database
