package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Ch1_Stdin() {

	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

}
