package main

import "time"
import "fmt"

func main() {
	//在本例中，假设我们执行了一个外部调用，它在两秒后将结果返回到通道c1上
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	//这里是用select实现超时。res:=<-c1等待一个结果，而<-Time.After等待超时1秒后发送一个值。
	//由于select将在有第一个准备就绪的接收时继续，我们会在操作超过允许的1秒时进入超时事件。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	//如果我们允许一个更长的超时时间3秒，则将能成功得到c2的值，并打印
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
