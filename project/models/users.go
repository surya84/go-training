package models

type User struct {
	Name     string
	Email    string
	Id       int
	Password string
	Mobile   int
	Gender   string

	Perms map[string]bool
}

// func User() {

// }
