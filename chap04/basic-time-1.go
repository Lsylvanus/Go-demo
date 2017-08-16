package main

import (
	"time"
	"fmt"
)

func main() {
	time.AfterFunc(5 * time.Second, f1)
	time.AfterFunc(2 * time.Second, f2)
	fmt.Println("main thread")
	time.Sleep(10 * time.Second)
}


func f1() {
	fmt.Println("f1 done !")
}

func f2() {
	fmt.Println("f2 done !")
}