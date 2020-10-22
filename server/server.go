package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

// Message a structure containing a message and its sender
type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	fmt.Println("Listening for connections...")
	for {
		conn, err := ln.Accept()
		if err == nil {
			conns <- conn
		}
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.

	clientMsg := bufio.NewReader(client)
	for {
		msg, err := clientMsg.ReadString('\n')
		if err == nil {
			msgs <- Message{clientid, fmt.Sprintf("Client %d> %s", clientid, msg)}
			fmt.Printf("Debug> Wrote message from client %d: %s", clientid, msg)
		}
	}

}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", *portPtr)
	fmt.Println("Debug> Listening to " + *portPtr + "...")

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)

	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
			fmt.Println("Debug> Found a new client!")
			clientID := len(clients)
			clients[clientID] = conn
			go handleClient(conn, clientID, msgs)

		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for clientID, client := range clients {
				fmt.Printf("Debug> Sending message to %d clients\n", len(clients))
				if clientID != msg.sender {
					fmt.Fprintf(client, msg.message)
				}
			}
		}
	}
}
