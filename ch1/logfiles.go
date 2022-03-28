package ch1

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "./test.log"

func Logging() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "Test@test ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags | log.Lshortfile) //log.Lshortfile - Номер строки
	iLog.Println("Запись лог")
	iLog.Println("Выход")
}
