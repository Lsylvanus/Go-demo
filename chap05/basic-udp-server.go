package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

func main() {
	//get a udp address
	udp_address, err := net.ResolveUDPAddr("udp4", ":9848")
	if err != nil {
		fmt.Println("Error happened when getting a udp address...", err)
		os.Exit(1)
	}

	//listen to the udp address
	udp_connection, err := net.ListenUDP("udp", udp_address)
	if err != nil {
		fmt.Println("Error happened when listening to the udp address...", err)
		os.Exit(1)
	}

	//greeting words
	str_header := "Hello, this is information from udp server , and current time is : "
	//main loop of udp server
	for {

		//read information from client
		var udp_buffer [1024]byte
		_, address, err := udp_connection.ReadFromUDP(udp_buffer[0:])
		if err != nil {
			continue
		}

		//get the current time infor
		str_time := time.Now().String()
		//set greeting words for client connection
		str_greetings := str_header + str_time

		//send information to client
		udp_connection.WriteToUDP([]byte(str_greetings), address)
	}
}