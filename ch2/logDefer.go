package ch1

import (
	"fmt"
	"log"
	"os"
	"time"
)

//Очень важная шутка(запомни)
//Логирование

//+ Логер с девефером
var LOGFILE = "./ch2/mGo.log"

func One(aLog *log.Logger) {
	aLog.Println("<-----------FUNCTION one ----------->")
	defer aLog.Println("<-----------FUNCTION one ----------->")
	for i := 0; i < 10; i++ {
		aLog.Println("Запись:", i)
	}
}

func Two(aLog *log.Logger) {
	aLog.Println("<-----------FUNCTION two----------->")
	defer aLog.Println("<-----------FUNCTION two----------->")
	for i := 10; i > 0; i-- {
		aLog.Println("Запись:", i)
	}
}

func Mainlog() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	start := time.Now()

	iLog := log.New(f, "logDefer ", log.LstdFlags)
	iLog.Println("<-----------Log start----------->")
	iLog.Printf("			     %s				   ", start.Format("Mon Jan 2 15:04:05 MST 2006"))

	One(iLog)
	Two(iLog)
	defer func() {
		end := time.Now()
		duratatio := time.Since(start)
		iLog.Printf("			     %s				   ", end.Format("Mon Jan 2 15:04:05 MST 2006"))
		iLog.Printf("			     %s				   ", duratatio)
		iLog.Println("<-----------Log end----------->")
	}()

}
