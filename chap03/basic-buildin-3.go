package main

import "fmt"

func main() {
	var c1 = complex(1.1, 2)
	var r1 = real(c1)
	var i1 = imag(c1)
	println("c1=", c1, " r1=", r1, " i1=", i1)
	fmt.Println("c1=", c1, " r1=", r1, " i1=", i1)
}
