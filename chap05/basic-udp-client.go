package main

import (
	"os"
	"fmt"
	"net"
)

func main() {

	//usage of client application
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	//get a udp address
	arg_address := os.Args[1]
	udp_address, err := net.ResolveUDPAddr("udp4", arg_address)
	if err != nil {
		fmt.Println("Error happened when getting a udp address...", err)
		os.Exit(1)
	}

	//connect to a udp server
	udp_connection, err := net.DialUDP("udp", nil, udp_address)
	if err != nil {
		fmt.Println("Error happened when connecting a udp server...", err)
		os.Exit(1)
	}

	//send information to udp server
	_, err = udp_connection.Write([]byte("hello"))
	if err != nil {
		fmt.Println("Error happened when sending message to the udp server...", err)
		os.Exit(1)
	}

	//udp buffer
	var udp_buffer [1024]byte

	//read information from udp server
	num, err := udp_connection.Read(udp_buffer[0:])
	if err != nil {
		fmt.Println("Error happened when reading from the tcp server...", err)
		os.Exit(1)
	}

	//display the information
	fmt.Println(string(udp_buffer[0:num]))
	os.Exit(0)
}