package main

import (
	"fmt"
	"go-training/sum"
)

func main() {
	sum.Add()
	sum.Sum = 10
	fmt.Println("sum:=", sum.Sum)

}
