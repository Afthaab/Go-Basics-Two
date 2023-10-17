package main

import (
	"net/http"
	"project-4/handler"
)

func main() {
	http.HandleFunc("/home", handler.HomeFunction)

	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		panic("Server not started")
	}
}
