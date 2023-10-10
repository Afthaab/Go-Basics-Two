package main

import (
	"fmt"
)

type user struct {
	name string
}

func main() {

	user1 := user{
		name: "AFThab",
	}
	checkType(user1)
	checkType(20)

}

func checkType(i any) {
	switch i.(type) {
	case int:
		fmt.Println("it is an int value")
	case user:
		fmt.Println("it is a struct")
	}
}
