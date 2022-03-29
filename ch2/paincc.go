package ch1

import (
	"fmt"
	"os"
)

// panic -- останавливает код
func Panicc() {
	if len(os.Args) == 1 {
		panic("Укажите аргументы")
	}
	fmt.Println("Спасибо что указали аргументы")

}
