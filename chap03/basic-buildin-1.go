package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["Mon"] = 1
	m1["Tue"] = 2
	m1["Wed"] = 3

	fmt.Println("before delete : m1=", m1, "len(m1)=", len(m1))

	println("delete element by using key Tue")
	delete(m1, "Tue")

	fmt.Println("after  delete : m1=", m1, "len(m1)=", len(m1))

}