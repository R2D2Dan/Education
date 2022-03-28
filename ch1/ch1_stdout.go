package ch1

import (
	"io"
	"os"
)

func Ch1_Stdout() {
	text := ""
	arguments := os.Args
	if len(arguments) == 1 {
		text = "Пожалуйста укажие аргумент"
	} else {
		text = arguments[1]
	}

	io.WriteString(os.Stdout, text)
	io.WriteString(os.Stdout, "\n")
}
