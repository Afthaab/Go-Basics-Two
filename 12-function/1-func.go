package main

import "fmt"

type operation func(x, y int) int

func main() {
	mul := func(a, b int) int {
		return a * b
	}
	Operate(mul, 20, 30)
	Operate(add, 20, 30)
	Operate(Sub, 20, 30)
}

func Operate(op operation, a, b int) {
	fmt.Println(op(a, b))
}

func add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}
