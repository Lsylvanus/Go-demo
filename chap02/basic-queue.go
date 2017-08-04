package main

import "fmt"

var queue []int

func main() {
	//fmt.Println("Hello World!")
	queue = make([]int, 0, 0)
	var inputCount int
	fmt.Scanln(&inputCount)

	var flag string
	var value int
	var queueLength int
	for i := 0; i < inputCount; i++ {
		fmt.Scanln(&flag, &value)
		queueLength = len(queue)
		if flag == "E" {
			queue = append(queue,value)
			queueLength = len(queue)
			fmt.Println(queueLength)
		}else if flag == "D"{
			if queueLength ==0 {
				fmt.Println("-1 0")
			}else{
				exitvalue := queue[0]
				queue = queue[1:]
				queueLength = len(queue)
				fmt.Println(exitvalue, queueLength)
			}
		}
	}
}