package postgres

import (
	"fmt"
	"project4/stores"
)

type Conn struct {
	db string
}

func (c Conn) Create(usr stores.User) error {
	fmt.Println("this is create method postgres ", usr)
	return nil
}

func (c Conn) Update(usr stores.User) error {
	fmt.Println("this is create postgres ", usr)
	return nil
}

func (c Conn) Delete(usr stores.User) error {
	fmt.Println("this is create postgres ", usr)
	return nil
}

func NewConn(db string) Conn {
	return Conn{
		db: db,
	}
}
