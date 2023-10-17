package main

import "fmt"

type Speaker interface {
	Speak()
}
type human struct {
	name string
}

type ai struct {
	name string
}

func (h human) Speak() {
	fmt.Println("human speaking", h.name)
}

func (a ai) Speak() {
	fmt.Println("ai speaking", a.name)
}

func doSomething(s Speaker) {
	fmt.Printf("%T\n", s)
	s.Speak()
}

func main() {
	u := human{
		name: "dev",
	}

	u2 := human{
		name: "alexa",
	}

	doSomething(u)
	doSomething(u2)
}
