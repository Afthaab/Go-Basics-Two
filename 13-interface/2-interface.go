package main

import (
	"fmt"
	"log"
)

type user struct {
	name string
}

// Write implements io.Writer.
func (u user) Write(p []byte) (n int, err error) {
	fmt.Println("Writer methond implemented")
	fmt.Println(u.name)
	return n, nil
}

func main() {
	u := user{name: "ajay"}
	l := log.New(u, "sales:", log.Lshortfile)
	l.Println()
}
