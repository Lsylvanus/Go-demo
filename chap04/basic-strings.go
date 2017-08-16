package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	fmt.Println(strings.HasPrefix(str, "He"))
	fmt.Println(strings.ContainsAny(str, "od"))
	fmt.Println(strings.Contains(str, "od"))
}