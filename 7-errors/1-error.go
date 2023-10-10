package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
	fmt.Println("Doing some work")
	fmt.Println("end of main")
}

func greet() {
	data := os.Args[1:]
	if len(data) != 3 {
		log.Println("will not work")
	}
	name := data[1]
	age := data[2]
	new, err := strconv.Atoi(name)
	// var i error
	if err != nil {
		return
	}
	fmt.Println(err, new)

	a, err := strconv.Atoi(age)
	if err != nil {
		return
	}
	fmt.Println(name, a)

}
