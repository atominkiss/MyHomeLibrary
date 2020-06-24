package main

import (
	. "MyHomeLibrary/Config"
	. "MyHomeLibrary/DAO"
	. "MyHomeLibrary/Model"

	"github.com/gorilla/mux"
	"labix.org/v2/mgo/bson"

	"encoding/json"
	"log"
	"net/http"
)

var config = Config{}
var booksDao = BooksDAO{}

func init() {
	config.Read()

	booksDao.Server = config.Server
	booksDao.Database = config.Database
	booksDao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil { // nolint:wsl
		log.Fatal(err)
	}
}

// GET list of books.
func getBooks(w http.ResponseWriter, _ *http.Request) {
	books, err := booksDao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, books)
}

// GET a book by its ID.
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := booksDao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book Id")
		return
	}
	respondWithJSON(w, http.StatusOK, book)

}

// POST a new book.
func createBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	book.ID = bson.NewObjectId()
	if err := booksDao.Insert(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, book)
}

// PUT update an existing book.
func updateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := booksDao.Update(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := booksDao.Delete(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}
