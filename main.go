package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// Book represents a book
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books []Book

func main() {
    // Initialize books
    books = append(books, Book{ID: "1", Title: "Book One", Author: "John Doe"})
    books = append(books, Book{ID: "2", Title: "Book Two", Author: "Jane Doe"})

    // Create a new router
    router := mux.NewRouter()

    // Route handlers
    router.HandleFunc("/books", getBooks).Methods("GET")
    router.HandleFunc("/books/{id}", getBook).Methods("GET")
    router.HandleFunc("/books", createBook).Methods("POST")
    router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

    fmt.Println("Starting server at port 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, book := range books {
        if book.ID == params["id"] {
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
    var newBook Book
    _ = json.NewDecoder(r.Body).Decode(&newBook)
    books = append(books, newBook)
    json.NewEncoder(w).Encode(newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, book := range books {
        if book.ID == params["id"] {
            var newBook Book
            _ = json.NewDecoder(r.Body).Decode(&newBook)
            books[index] = Book{
                ID:     params["id"],
                Title:  newBook.Title,
                Author: newBook.Author,
            }
            json.NewEncoder(w).Encode(books[index])
            return
        }
    }
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, book := range books {
        if book.ID == params["id"] {
            books = append(books[:index], books[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(books)
}
