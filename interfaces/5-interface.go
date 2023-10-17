package main

import (
	"fmt"
)

type Human struct {
	name string
}

type walker interface {
	Walk()
}
type runner interface {
	Run()
}

type walkRunner interface {
	walker
	runner
}

func (h Human) Walk() {
	fmt.Println(h.name, "walking")
}

func (h Human) Run() {
	fmt.Println(h.name, "running")
}

func main() {
	a := Human{
		name: "ade",
	}

	var w walker = a
	var r runner = a
	var rw walkRunner = a
	fmt.Println(w)
	fmt.Println(r)
	rw.Run()
	rw.Walk()

}
