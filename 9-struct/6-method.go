package main

import "fmt"

type music struct {
	lang  string
	title string
}

type director struct {
	dname string
}

type authour struct {
	name string
	music
	director
}

func (t *music) update(tit string) {
	t.title = tit
}

func main() {
	a := authour{
		name: "dhoni",
		music: music{
			lang:  "telugu",
			title: "jab thak",
		},

		director: director{"abfd"},
	}

	fmt.Println(a.lang)
	fmt.Println(a.name)
	fmt.Println(a.title)

	a.update("anthem")

	fmt.Println(a.title)
	fmt.Println(a.dname)
}
