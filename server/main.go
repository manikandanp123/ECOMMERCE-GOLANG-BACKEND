package main

import (
	"fmt"
	"golang-backend/routes"
	"log"
	"net/http"
)

func main() {
	route := routes.Router()
	fmt.Println("server is at port 8080")
	log.Fatal(http.ListenAndServe(":8080", route))
}
