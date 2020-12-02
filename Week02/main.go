package main

import (
	"github.com/x-debug/Go-000/Week02/api"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users", api.UserHandler)

	log.Fatalln(http.ListenAndServe(":1234", nil))
}
