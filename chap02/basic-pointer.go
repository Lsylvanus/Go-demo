package main

import "fmt"

func main() {
	var intPointer *int
	var i int = 100
	intPointer = &i

	fmt.Printf("value of i = [%d] address is [%x] \n", i, &i)
	fmt.Println("value of intPointer = ", intPointer)
	fmt.Println("value of var pointed to = ", *intPointer)
}