package main

import (
	"os"
	"net"
	"fmt"
	"io/ioutil"
)

func main() {

	//usage of client application
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	//get a tcp address
	arg_address := os.Args[1]
	tcp_address, err := net.ResolveTCPAddr("tcp4", arg_address)
	if err != nil {
		fmt.Println("Error happened when getting a tcp address...", err)
		os.Exit(1)
	}

	//connect to a tcp server
	tcp_connection, err := net.DialTCP("tcp", nil, tcp_address)
	if err != nil {
		fmt.Println("Error happened when connecting a tcp server...", err)
		os.Exit(1)
	}

	//read information from tcp server
	str_greeting, err := ioutil.ReadAll(tcp_connection)
	if err != nil {
		fmt.Println("Error happened when reading from the tcp server...", err)
		os.Exit(1)
	}

	//display the information
	fmt.Println(string(str_greeting))
	os.Exit(0)
}