package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

func main() {
	//get a tcp address
	tcp_address, err := net.ResolveTCPAddr("tcp4", ":8848")
	if err != nil {
		fmt.Println("Error happened when getting a tcp address...", err)
		os.Exit(1)
	}

	//listen to the tcp address
	tcp_listener, err := net.ListenTCP("tcp", tcp_address)
	if err != nil {
		fmt.Println("Error happened when listening to the tcp address...", err)
		os.Exit(1)
	}

	//greeting words
	str_header := "Hello, this is information from tcp server , and current time is : "
	//main loop of tcp server
	for {

		//waiting connection from client
		tcp_connection, err := tcp_listener.Accept()
		if err != nil {
			continue
		}

		//get the current time infor
		str_time := time.Now().String()
		//set greeting words for client connection
		str_greetings := str_header + str_time

		//send information to client
		tcp_connection.Write([]byte(str_greetings))
		tcp_connection.Close()
	}
}