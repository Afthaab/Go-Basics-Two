package models

import "errors"

var users = map[uint64]User{
	123: {
		FName:    "Bob",
		LName:    "abc",
		Password: "someSecretPassword",
		Email:    "bob@email.com",
	},
}

var ErrNotFound = errors.New("the value is not found")

func FetchUser(userId uint64) (User, error) {
	userData, ok := users[userId]
	if !ok {
		return User{}, ErrNotFound
	}
	return userData, nil
}
