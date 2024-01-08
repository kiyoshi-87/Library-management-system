package main

import (
	"fmt"
	"net/http"

	"github.com/kiyoshi-87/library-management-system/PKG/routes"
)

func main() {
	fmt.Println("Connecting mySQL to our project!")
	r := routes.Router()

	http.Handle("/", r)
	fmt.Println("Starting the server on port 8000...")
	if err := http.ListenAndServe(":8000", r); err != nil {
		panic(err)
	}
}
