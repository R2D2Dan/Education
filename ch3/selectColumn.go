package ch1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func SelectColumnMain() {
	argument := os.Args
	if len(argument) < 2 {
		fmt.Println("usage: selectColumn column <file1> [<file2> [... <fileN]]\n")
		os.Exit(1)
	}

	temp, err := strconv.Atoi(argument[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	column := temp
	if column < 0 {
		fmt.Println("Ошибка столбца")
		return
	}

	for _, filename := range argument[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("Ошбка чтения файла", err, filename)
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)

		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}

	}

}

//go run main.go 4  ./ch2/mGo.log
