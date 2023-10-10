package main

import (
	"fmt"
)

var user = make(map[int][]string)

type money int //int is a underlying type and money is a new type

func main() {
	// arr := []string{"afthab"}
	// afthab := []int{10, 20, 30, 40, 50}
	// user := make(map[[5]int]string)

	// user[1] = arr
	// user[2] = arr

	var rupee money = 10

	var i int64

	// i := 120

	rupee = money(i)

	fmt.Printf("%T", rupee)
	// search(2)

}

func search(id int) {
	if _, ok := user[id]; !ok {
		fmt.Println("it is not present")
		return
	}
	fmt.Println("It is present")
}
