package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.NewTimer(10 * time.Second)
	//v := <- t.C
	//fmt.Println(v)
	go onTime(t.C)
	fmt.Println("main thread")
	time.Sleep(10 * time.Second)

}

func onTime(c <-chan time.Time) {
	for now := range c {
		// now := <- c
		fmt.Println("onTime", now)
	}
}