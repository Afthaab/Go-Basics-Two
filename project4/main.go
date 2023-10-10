package main

import (
	"project4/stores"
	"project4/stores/mysql"
)

func main() {
	user := stores.User{
		Name:  "Mohammed Afthab",
		Email: "afthab12@gmail.com",
	}
	db := mysql.NewConn("mysql")
	// stores.StorerInterface = db
	// db.Create(user)

	// pb := postgres.NewConn("postgres")
	// stores.StorerInterface = pb
	// stores.StorerInterface.Create(user)
	// pb.Create(user)

	newS := stores.NewService(db)
	newS.Update(user)
}
