package main

import "fmt"

type Caller interface {
	Info()
}

type Zap struct {
	path string
}
type Logrus struct {
	path string
}

func (z Zap) Info() {
	fmt.Println("in Zap ", z.path)
}

func (l Logrus) Info() {
	fmt.Println("in logrus ", l.path)
}

func doLog(f Caller) {
	f.Info()
}

func main() {
	u := Zap{
		path: "Terminal",
	}

	u2 := Logrus{
		path: "File",
	}

	doLog(u)
	doLog(u2)
	u.Info()
	u2.Info()

	// fmt.Println(u)
	// fmt.Println(u2)
}
