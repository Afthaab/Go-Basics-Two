package main

import (
	"fmt"
)

func main() {
	doggy := dog{
		name: "poorvi",
	}

	rw = doggy

	rw.walk()
	rw.run()

}

type dog struct {
	name string
}

type read interface {
	walk()
}

type write interface {
	run()
}

type readwriter interface {
	read
	write
}

var rw readwriter

func (d dog) walk() {
	fmt.Println(d.name, "walk")
}

func (d dog) run() {
	fmt.Println(d.name, "run")
}
