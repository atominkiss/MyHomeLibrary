package dao

import (
	. "MyHomeLibrary/Model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type BooksDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const COLLECTION = "books"

// Establish a connection to database
func (m *BooksDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of books
func (m *BooksDAO) FindAll() ([]Book, error) {
	var books []Book
	err := db.C(COLLECTION).Find(bson.M{}).All(&books)
	return books, err
}

// Find a book by its id
func (m *BooksDAO) FindById(id string) (Book, error) {
	var book Book
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

// Insert a book into database
func (m *BooksDAO) Insert(book Book) error {
	err := db.C(COLLECTION).Insert(&book)
	return err
}

// Delete an existing book
func (m *BooksDAO) Delete(book Book) error {
	err := db.C(COLLECTION).Remove(&book)
	return err
}

// Update an existing book
func (m *BooksDAO) Update(book Book) error {
	err := db.C(COLLECTION).UpdateId(book.ID, &book)
	return err
}
