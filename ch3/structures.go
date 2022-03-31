package ch1

import (
	"fmt"
)

type XYZ struct {
	X int
	Y int
	Z int
}

func StructuresMain() {

	test := XYZ{33, 55, 44}

	fmt.Println(test.X, test.Y, test.Z)

	pSlice := [4]XYZ{}
	fmt.Println(pSlice)
	pSlice[0].X = 2
	pSlice[0].Y = 3
	pSlice[0].Z = 1
	fmt.Println(pSlice)

}
