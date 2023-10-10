package main

import (
	"firstproject/models"
	"fmt"
)

func main() {
	userOne := models.User{
		Name:     "Mohammed Afthab",
		Email:    "mohammedafthab15@gmail.com",
		Password: "qwert12345",
		Permission: map[string]bool{
			"user": false,
		},
		Hobbies: []string{"Eating, Drinking"},
	}
	userTwo := models.User{
		Name:     "Lione Messi",
		Email:    "messi10@gmail.com",
		Password: "messi_is_the_best",
		Permission: map[string]bool{
			"admin": true,
		},
		Hobbies: []string{"Movies"},
	}

	fmt.Println(userOne)
	fmt.Println(userTwo)

}
