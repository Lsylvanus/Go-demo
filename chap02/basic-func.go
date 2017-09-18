package main

import "fmt"

func a() (int, int) {
	return 1, 3
}

func b() (int, int) {
	return 2, 4
}

func main() {
	x := []func() (int, int){a, b}
	for _, v := range x {
		a, b := v()
		fmt.Printf("values is %d and %d.\n", a, b)
	}
	s()
}

func s() {
	b := true
	var a interface{}
	a = b
	switch a.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Printf("unexpected type %T, its value is : %v\n", a, a)
	}
}