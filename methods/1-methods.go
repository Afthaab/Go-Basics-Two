package main

import "fmt"

type user struct {
	name     string
	password string
}

var db = make(map[int]user)

func (u *user) Add(id int) {
	db[id] = *u
}

func (u *user) Update(name string, id int) {
	u.name = name
	db[id] = *u
}

func main() {
	// userDetails1 := user{
	// 	name:     "Mohammed Afthab",
	// 	password: " afthu",
	// }
	// userDetails2 := user{
	// 	name:     "Lionel Messi",
	// 	password: " 10001",
	// }
	// userDetails1.Add(1)
	// userDetails2.Add(2)

	// fmt.Println(db)

	// userDetails1.Update("Ronaldo", 1)

	// fmt.Println(db)

	userDetails := []user{
		{
			name:     "Mohammed Afthab",
			password: "qwert",
		},
		{
			name:     "Tejaswinin",
			password: "1234",
		},
		{
			name:     "poorvi",
			password: "122345",
		},
		{
			name:     "jeevan",
			password: "poiiyy",
		},
	}

	for i, v := range userDetails {
		fmt.Println(i, v)
	}
	fmt.Println(userDetails)

}
