package main

import "fmt"

func far(from string) {
	for i := 0; i < 30; i++ {
		fmt.Println(from, ";", i)
	}
}

func main() {
	//假设我们有个f(s)的函数调用。这里我们通过一般方法调用，令其同步执行
	far("direct")
	//要想让这个函数在goroutine中触发，使用go f(s)。这个新的goroutine将会与调用它的并行执行
	go far("goroutine")
	//我们也可以启动一个调用匿名函数的goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	//现在，这两个方法调用在独立的goroutine中异步执行了，故方法执行直接落到了这里
	/*
	goroutine ; 0
	goroutine ; 1
	going
	goroutine ; 2
	goroutine ; 3
	goroutine ; 4
	*/
	//Scanln代码需要在程序退出前按下一个键
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
