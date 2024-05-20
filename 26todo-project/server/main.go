package main

import (
	"fmt"
	"net/http"

	"github.com/atharv1809/golang-todo-app/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on Port: 9000")
	http.ListenAndServe(":9000", r)
}
