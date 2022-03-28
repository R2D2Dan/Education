package ch1

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Errormain() {

	if len(os.Args) == 1 {
		fmt.Println("Укажите аргументы")
		os.Exit(1)
	}

	argumnet := os.Args
	//Создаем новую ошибку
	var err error = errors.New("Error")
	k := 1
	var n float64
	for err != nil {
		if k > len(argumnet) {
			fmt.Println("Нет аргументов")
			return
		}

		n, err = strconv.ParseFloat(argumnet[k], 64)
		if err != nil {
			log.Println("Ошибка аргумента")
		}
		k++
	}
	min, max := n, n
	for i := 2; i < len(argumnet); i++ {
		n, err = strconv.ParseFloat(argumnet[i], 64)
		if err != nil {
			log.Println("Ошибка аргумента")
			return
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}

	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

}
