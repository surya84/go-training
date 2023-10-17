// package main

// import "errors"

// type Student struct {
// 	Name  string `json:"S_Name"`
// 	Id    int    `json:"S_Id"`
// 	Age   int    `json:"S_Age"`
// 	Grade int    `json:"S-Grade"`
// }

// var StudentData = map[uint64]Student{
// 	184: {
// 		Name:  "surya",
// 		Id:    48,
// 		Age:   22,
// 		Grade: 80,
// 	},

// 	146: {
// 		Name:  "teja",
// 		Id:    45,
// 		Age:   23,
// 		Grade: 82,
// 	},
// }

// var ErrUserNotFound = errors.New("User Not Found")

// func GetUser(userId uint64) (User, error) {
// 	user, err := userMap[userId]

// 	if !err {
// 		return User{}, ErrUserNotFound
// 	}

// 	return user, nil
// }
