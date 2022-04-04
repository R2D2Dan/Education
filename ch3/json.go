package ch1

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}
type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func SaveJson(filename *os.File, key interface{}) bool {
	encodingJS := json.NewEncoder(filename)
	err := encodingJS.Encode(key)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

func MainJSON() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}
	filename := arguments[1]
	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord.Name)
		fmt.Println(myRecord.Tel)
	} else {
		fmt.Println(err)
	}

	MyRecord1 := Record{
		Name:    "Stive",
		Surname: "Andres",
		Tel: []Telephone{
			Telephone{Mobile: true, Number: "231564989"},
			Telephone{Mobile: false, Number: "asdw23aw"},
		},
	}

	SaveJson(os.Stdout, MyRecord1)

}
