package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args[1:]
	// fmt.Printf("%T", os.Args)
	fmt.Println(a)
	// fmt.Println(os.Args)
}
