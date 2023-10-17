package main

import "fmt"

type Student struct {
	Name    string
	Age     int
	Degreee string
}

type Student1 struct {
	Name string
	Age  int
	stu  Student
}

func (s *Student1) print() {
	s.Name = "poorvi"
	fmt.Println(s.Name)
}

func main() {
	afthab := Student1{
		Name: "Afthab",
		Age:  22,
		stu: Student{
			Name:    "MESSIE",
			Age:     46,
			Degreee: "BTECH",
		},
	}
	afthab.print()
	fmt.Println(afthab)

}
