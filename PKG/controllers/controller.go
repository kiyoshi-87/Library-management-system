package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kiyoshi-87/library-management-system/PKG/models"
)

//========> DB helper Functions <========

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//SLice to return
	var books []models.Book

	//Raw query to display to select all the books in the table
	models.DB.Find(&books)

	//Return the slice
	json.NewEncoder(w).Encode(books)
}

// func Home(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Library Management system API is running!"))
// }

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookID := params["id"]

	// Convert bookID to uint
	id, err := strconv.Atoi(bookID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid book ID!")
	}

	var book models.Book

	result := models.DB.First(&book, id)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No book found with that ID")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	models.DB.Create(&book) //CREATE
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookID := params["id"]

	id, err := strconv.Atoi(bookID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid book ID!")
	}

	var book models.Book

	result := models.DB.First(&book, id)

	if result.Error != nil {
		json.NewEncoder(w).Encode("No book found with that ID")
	}
	models.DB.Delete(&book) //DELETE
	json.NewEncoder(w).Encode("Book deleted")
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	models.DB.Delete(&books)
	json.NewEncoder(w).Encode("All books deleted")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookID := params["id"]

	id, err := strconv.Atoi(bookID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid book ID!")
	}

	var book models.Book

	result := models.DB.First(&book, id)

	if result.Error != nil {
		json.NewEncoder(w).Encode("No book found with that ID")
	}
	json.NewDecoder(r.Body).Decode(&book)
	models.DB.Save(&book) //UPDATE=SAVE
	json.NewEncoder(w).Encode(book)
}
