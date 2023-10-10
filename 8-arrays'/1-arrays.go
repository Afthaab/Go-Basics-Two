package main

import (
	"fmt"
)

func main() {
	// var arr [5]int
	// fmt.Println(arr)
	// b := [5]int{10, 5, 4, 6, 3}
	// fmt.Println(b)
	// c := [...]int{100, 200}
	// fmt.Println(c)
	// Slices()
	// NewSlice()
	NewMake()
}

func Newfunc()  {
	
}

func NewMake() {
	a := []int{10, 20, 30, 40, 50}
	b := make([]int, len(a[1:3]))
	copy(b, a[1:3])
	fmt.Println(a, b)
}

func Slices() {
	// a := []int{10, 20, 30, 40, 50, 60}

	// b := a[2:5]

	// b = append(b, 999)
	// a = append(a, 90)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(cap(a))
	// fmt.Println(len(a))

	// var arr []int

	// sls := []int{}

	// fmt.Println(cap(sls))
	// fmt.Println(len(sls))

	i := []int{10, 20, 30, 40, 50, 70}

	b := i[1:5]

	b[0] = 99
	b = append(b, 60)

	fmt.Println(i, b)
	fmt.Print(cap(i))
}

func NewSlice() {
	a := make([]int, 0, 5)

	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	fmt.Printf("%p", a)
	a = append(a, 100, 300, 400, 500, 500, 500, 700)
	fmt.Println("/////////////")
	fmt.Printf("%p", a)
	fmt.Println(cap(a))
	fmt.Println(len(a))

	// b := []int{}
	// fmt.Println(len(b))
	// fmt.Println(cap(b))
	// fmt.Println(b)
}
