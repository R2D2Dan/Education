package ch1

import (
	"errors"
	"log"
)

func NewError(a, b int) error {
	if a == b {
		err := errors.New("error in NewErorr() function")
		return err
	} else {
		return nil
	}

}

func ErrorMain() {
	err := NewError(1, 2)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(err)

	err = NewError(1, 1)
	if err != nil {
		log.Println(err)
	}

	if err.Error() == "Error in NewErorr() function!" {
		log.Println(err)
		return
	}

}
