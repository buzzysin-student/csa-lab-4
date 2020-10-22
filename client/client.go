package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

const (
	eol = '\n'
)

func read(conn *net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	response := bufio.NewReader(*conn)
	for {
		msg, err := response.ReadString(eol)
		if err == nil {
			fmt.Println(msg)
		}
	}
}

func write(conn *net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	input := bufio.NewReader(os.Stdin)
	for {
		msg, _ := input.ReadString(eol)
		fmt.Fprintf(*conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", ":8030", "IP:port string to connect to")
	flag.Parse()

	//TODO Try to connect to the server
	conn, err := net.Dial("tcp", *addrPtr)
	if err != nil {
		panic("No server is running!")
	} else {
		fmt.Println("Found server running at " + *addrPtr + "...")
	}

	//TODO Start asynchronously reading and displaying messages
	go read(&conn)
	//TODO Start getting and sending user messages.
	for {
		write(&conn)
	}
}
