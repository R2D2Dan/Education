package ch1

import (
	"fmt"
	"runtime"
)

type data struct {
	i, j int
}

//Руководство как засрать себе память
func SliceGC() {
	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}
	runtime.GC()
	_ = structure[0]
	fmt.Println(structure)
}

func MapStar() {
	var N = 40000000
	myMap := make(map[int]*int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = &value
	}
	runtime.GC()
	_ = myMap[0]
	fmt.Println(myMap)
}

func MapNoStar() {
	var N = 40000000
	myMap := make(map[int]int)

	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = value
	}
	runtime.GC()
	_ = myMap[0]
	fmt.Println(myMap)
}
