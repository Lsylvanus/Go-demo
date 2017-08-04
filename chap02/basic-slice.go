package main

import "fmt"

func main() {

	//def & init 1
	s1 := []float32{1.1, 2.2, 3, 4, 5}
	fmt.Println("s1 = ", s1)

	//def & init 2
	s2 := s1[2:4]
	fmt.Println("s2 = ", s2)

	//def & init 3
	s3 := s1[:4]
	fmt.Println("s3 = ", s3)

}