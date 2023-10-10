package main

import "fmt"

type Speaker interface {
	Speak()
}
type human struct {
	name string
}

type ai struct {
	name string
}

func (h human) Speak() {
	fmt.Println("This is human speaking", h.name)

}

func (a ai) Speak() {
	fmt.Println("This is alexa speaking", a.name)
}

func doSomething(s Speaker) {
	s.Speak()
}

// func main() {
// 	afthab := human{
// 		name: "Afthab",
// 	}
// 	doSomething(afthab)
// }
