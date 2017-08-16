package main

import (
	"time"
	"fmt"
)

var count int = 0

func main() {
	t := time.Tick(2 * time.Second)

	i := 0
	for now := range t {
		fmt.Println(now, doSomething())
		i++
		if i > 10 {
			break
		}
	}
}

func doSomething() int {
	count++
	return count
}