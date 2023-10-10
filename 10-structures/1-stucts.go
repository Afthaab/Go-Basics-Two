package main

import "fmt"

type user struct {
	name   string
	age    int
	degree string
	marks  []int
}

func main() {
	// mark := []int{100, 99, 89, 99, 100}
	userDetails := user{
		name:   "Mohammed Afthab",
		age:    22,
		degree: "BCA",
		marks:  []int{10, 20, 30, 40},
	}

	fmt.Printf("%+v \n", userDetails)
	fmt.Printf("%#v \n", userDetails)
}
