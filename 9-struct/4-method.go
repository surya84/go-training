package main

import "fmt"

type Subject struct {
	LangName string
}

func (p *Subject) update(lang string) {
	p.LangName = lang
}

func main() {
	u1 := Subject{
		LangName: "english",
	}
	u1.update("maths")

	fmt.Println(u1)
}
