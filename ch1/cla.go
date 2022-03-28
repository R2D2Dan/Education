package ch1

import (
	"fmt"
	"os"
	"strconv"
)

//go run main.go 33 110 521 91 0 12
//min: 0
//max: 521
func Cla() {
	if len(os.Args) == 1 {
		fmt.Println("Пожалуйста укажите аргумент строки")
		os.Exit(1)
	}
	arguments := os.Args

	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("min:", min)
	fmt.Println("max:", max)

}
