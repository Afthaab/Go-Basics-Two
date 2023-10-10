package main

import "fmt"

func key() int {
	return 10
}

func main() {
	if a := key(); a > 0 { // variable declared in the short if will get expired after the if
		fmt.Println("true if")
	}
	if a, _ := fmt.Println("hello"); a > 5 {
		fmt.Println("true if again", a)
	}
}
