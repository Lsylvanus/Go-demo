package main

import "fmt"
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
}
