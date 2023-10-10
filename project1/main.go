package main

import (
	"fmt"
	"project2/db"
)

func main() {
	dbConnection, ok := db.NewConn("aftah")
	if !ok {
		fmt.Println("String is empty")
		return
	}

	fmt.Println(dbConnection)
	fmt.Println(db.Conn{})
}
