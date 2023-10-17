package main

import (
	"fmt"
)

// nil means no panic // otherwise msg var would be the msg of the panic
// recover func // it recovers from the panic and stop panic further propagation
func main() {

	doSomething()
	//panic("panic happens")
	fmt.Println("end of the main, stopped gracefully")

}

func doSomething() {
	defer recoveryFunc()
	// var i = []int{1, 2}

	var i []int
	i[1] = 1000
}

func recoveryFunc() {
	msg := recover()
	if msg != nil {
		//fmt.Println(string(debug.Stack()))
		fmt.Println(msg)
		fmt.Println("recovered from the panic")
	}
}
