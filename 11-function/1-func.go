package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(Greet("Afthab"))
	sum, err := add("a", "2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The sum : ", sum)
}

func add(a, b string) (int, error) {
	num, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	num1, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}
	return num + num1, nil
}

// func Greet(name string) string {
// 	return "Hello " + name
// }
