package main

import "fmt"

type User struct {
	Name   string
	Id     int
	Passwd string
}

var addMap = make(map[int]User)

func (u User) addUser(id int) {
	addMap[id] = u
}

func main() {
	u1 := User{
		Name:   "surya",
		Id:     1,
		Passwd: "asdf",
	}

	u2 := User{
		Name:   "teja",
		Id:     4,
		Passwd: "dsf",
	}

	u1.addUser(u1.Id)
	u2.addUser(u2.Id)
	for k, v := range addMap {
		fmt.Println(k, " : ", v)
	}

}
