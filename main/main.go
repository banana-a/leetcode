package main

import "fmt"

func main() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	test1 := append(test, 100000)
	test[0] = 1000
	fmt.Println(test)
	fmt.Println(test1)
}
