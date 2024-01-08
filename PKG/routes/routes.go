package routes

import (
	"github.com/gorilla/mux"
	"github.com/kiyoshi-87/library-management-system/PKG/controllers"
)

/*  ALL ROUTES

GET All books ---> /books
GET One book ---> /books/:id
POST One book ---> /books
Delete one Book ---> /books/:id
Update one Book ---> /books/:id

*/

func Router() *mux.Router {
	r := mux.NewRouter()

	//Handling the routes
	// r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("books/", controllers.AddBook).Methods("POST")
	r.HandleFunc("books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("books/{id}", controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc("books/", controllers.DeleteBooks).Methods("DELETE")

	return r
}
