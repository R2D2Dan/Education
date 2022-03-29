package ch1

import "fmt"

// defer во всей красе)
//210
func Deferr() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")

	}

}

//000
func Deferr2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

//123
func Deferr3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}
