package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", Home)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Println(err)
		return
	}
	// mux is your request matcher (multiplexer)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello this is our world"))
}
