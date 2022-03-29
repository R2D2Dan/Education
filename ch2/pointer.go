package ch1

import "fmt"

func getPointer(n *int) {
	*n = *n * *n
}
func returnPointer(n int) *int {
	v := n * n
	return &v
}
func MainPointer() {
	i := -10
	j := 25
	pI := &i
	pJ := &j
	fmt.Println("pI memory:", pI)
	fmt.Println("pI memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pI value:", *pJ)
	*pI = 123456
	*pI--
	fmt.Println("i:", i)
	getPointer(pJ)
	fmt.Println("j:", j)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
}
