package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	FirstName string `json:"name"`
	LastName  string `json:"lastname"`
}

var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get book with id 
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range books {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatalln("something went wrong with decode json")
	}
	book.ID = strconv.Itoa(rand.Intn(100))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range books {
		if item.ID == param["id"] {
			books = append(books[:index], books[index+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init books
	// @TODO: Using database 
	books = append(books, Book{ID: "1", Isbn: "124864", Title: "book 1", Author: &Author{FirstName: "john", LastName: "steven"}})
	books = append(books, Book{ID: "2", Isbn: "548618", Title: "book 2", Author: &Author{FirstName: "dave", LastName: "dao"}})
	books = append(books, Book{ID: "3", Isbn: "355482", Title: "book 3", Author: &Author{FirstName: "sara", LastName: "ross"}})

	//init router
	r := mux.NewRouter()

	// Router Handlers (URLs)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/book", createBook).Methods("POST")
	// r.HandleFunc("api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run Server on port :8000
	log.Fatal(http.ListenAndServe(":8000", r))

}
