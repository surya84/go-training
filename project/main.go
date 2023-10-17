package main

import (
	"fmt"
	"project/models"
)

func main() {
	u := models.User{
		Name:     "surya",
		Email:    "suryateja@gmail.com",
		Id:       24,
		Password: "surya!23",
		Mobile:   1234567890,
		Gender:   "Male",
		Perms: map[string]bool{
			"admin": true,
		},
	}

	fmt.Println("Name", " : ", u.Name)
	fmt.Println("Email", " : ", u.Email)
	fmt.Println("Id", " : ", u.Id)
	fmt.Println("Password", " : ", u.Password)
	fmt.Println("Mobile number", " : ", u.Mobile)
	fmt.Println("Gender", " : ", u.Gender)

	for k, v := range u.Perms {
		fmt.Println(k, " : ", v)
	}
	// fmt.Println(u.Perms)
}
