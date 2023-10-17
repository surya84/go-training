package models

import (
	"errors"
)

var users = map[uint64]User{
	123: {
		FName:    "Bob",
		LName:    "abc",
		Password: "someSecretPassword",
		Email:    "bob@email.com",
	},
}

var ErrUserIdNotFound = errors.New("user id not found")

func FetchUser(userId uint64) (User, error) {

	u, ok := users[userId]

	if !ok {
		return User{}, ErrUserIdNotFound
	}

	return u, nil

}
