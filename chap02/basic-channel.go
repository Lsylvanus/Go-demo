package main

import (
	"fmt"
	"time"
)

/**
通道是go所特有的数据类型之一，
详细地使用方法需要结合通道的长度/容量/单向或双向等才能更好地理解的这样一种通信手段
*/
func main() {
	fmt.Println("function begins ... ")
	c := make(chan bool)
	go func() {
		fmt.Println("func has been called.")
		close(c)
	}()
	<-c
	fmt.Println("Completed.")

	//通过make(chan val-type)创建新的channel
	//channel的类型依赖于它们要传递的值
	messages := make(chan string)
	//向channel传递值使用channel <- 语法。在这里我们从一个新的goroutine中发送了一个"ping"到message通道中
	go func() { messages <- "ping" }()
	//<-channel语法从channel中获取值。这里我们接收了上面发送的"ping"信息并打印
	msg := <-messages
	fmt.Println(msg)

	//这里我们创建了一个能够缓冲2个字符串值的channel
	ms := make(chan string, 2)
	//由于channel带有缓冲区，我们可以发送值，无需响应的并发接收
	ms <- "buffered"
	ms <- "channel"
	//稍后，我们像往常一样，接收了这两个值
	fmt.Println(<-ms)
	fmt.Println(<-ms)

	//启动一个worker goroutine，赋予它用以通知的channel
	done := make(chan bool, 1)
	go worker(done)
	//在channel接收到来自worker的通知前，保持阻塞
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

//这个方法将在一个goroutine中运行。
//done channel用来通知其他的goroutine这个方法执行完毕
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	//发送一个值来通知这里已经做完
	done <- true
}

//ping函数只接受一个发送数据的channel，如果试图在其上获取数据，将会引发编译时异常
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//pong函数接受一个通道用于接收(pings)，另一个用于发送(pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
