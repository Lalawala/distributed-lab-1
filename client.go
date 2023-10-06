package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	for {
		fmt.Println("Enter message:")
		message := ""
		fmt.Scanln(&message)
		conn.Write([]byte(strings.TrimSpace(message) + "\n"))
	}

}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	conn, err := net.Dial("tcp", *addrPtr)

	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
