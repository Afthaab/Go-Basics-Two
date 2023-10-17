package main

import (
	"net/http"
	"project5/handler"
)

func main() {

	http.HandleFunc("/get/user", handler.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
