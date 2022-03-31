package ch1

import "fmt"

type MyStr struct {
	Name     string
	LastName string
	Height   int
}

func NewStruct(n, l string, h int) *MyStr {
	if h > 300 {
		h = 0
	}
	return &MyStr{n, l, h}
}

func NewStructNotPointer(n, l string, h int) MyStr {
	if h > 300 {
		h = 0
	}
	return MyStr{n, l, h}
}

func NNew() {
	pS := new(MyStr)
	sP := new([]MyStr)
	fmt.Println("pS:", pS)
	fmt.Println("sP:", sP)
	pS.Name = "Test"
	*sP = make([]MyStr, 10)

	fmt.Println("pS:", pS)
	fmt.Println("*sP:", *sP)
	fmt.Println("sP:", sP)
}

func PointerStructMain() {
	s1 := NewStruct("Artem", "Testivich", 150)
	s2 := NewStructNotPointer("Artem", "Testivich", 150)
	fmt.Println(s1.Name)
	fmt.Println(s2.Name)
	s1.Height = 200
	fmt.Println(s1.Height)
}
