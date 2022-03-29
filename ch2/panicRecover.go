package ch1

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}
func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func MainPanicRecover() {
	a()
	fmt.Println("main() ended!")
}

/*
input:
	Inside a()
	About to call b()
	Inside b()
	Recover inside a()!
	main() ended!

*/
