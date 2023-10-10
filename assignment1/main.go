package main

import "fmt"

type req func(endpoint string)

func main() {
	request(get, "google")
	request(post, "youtube")
	request(put, "linkedIn")
}

func request(op req, name string) {
	op(name)
}

func get(endpoint string) {
	fmt.Println(endpoint, "get is called")
}

func post(endpoint string) {
	fmt.Println(endpoint, "post is called")
}

func put(endpoint string) {
	fmt.Println(endpoint, "put is called")
}
