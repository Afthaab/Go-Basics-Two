package main

import "fmt"

type user struct {
	name string
	age  string
}

type movie struct {
	moviename string
	user
	books
	perms
}

type books struct {
	bookname string
}

type perms struct {
	pername string
}

func main() {
	details := movie{
		moviename: "Kannur Squad",
		user: user{
			name: "Mammoty",
			age:  "74",
		},
		books: books{
			bookname: "dead of the devil",
		},
		perms: perms{
			pername: "poorvi",
		},
	}
	details.updateAge("65")
	fmt.Println(details)
}

func (us *movie) updateAge(age string) {
	us.age = age
}
