package ch1

import "fmt"

func retThree(n int) (n1 int, n2 int, n3 int) {
	return n * 2, n * n, -n

}

func TupelsMain() {
	fmt.Println(retThree(10))
	n1, n2, n3 := retThree(10)
	fmt.Println(n1, n2, n3)

	n1, n2 = n2, n1
	fmt.Println(n1, n2)
}
