package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Command to execute: go run client.go <server ip> <server port>")
		fmt.Println("Please pass server & port to connect to server. Exiting")
		os.Exit(0)
	}

	sendRequest(args[0], args[1])
}

func sendRequest(serverIP string, serverPort string) {
	conn, err := net.Dial("tcp", serverIP+":"+serverPort)
	if err != nil {
		fmt.Println("Problem connecting server at ", serverIP, serverPort, err.Error())
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go clientReceiver(conn, wg)
	go clientSender(conn, wg)

	wg.Wait()
	conn.Close()
}

func clientSender(conn net.Conn, wg sync.WaitGroup) {
	defer wg.Done()

	for {
		reader := bufio.NewReader(os.Stdin)
		txt, _ := reader.ReadString('\n')
		// fmt.Print("C_You: ")
		fmt.Fprintf(conn, txt)

		if strings.TrimSpace(string(txt)) == "STOP" {
			fmt.Println("Client exiting...")
			fmt.Fprintf(conn, txt)
			return
		}
	}
}

func clientReceiver(conn net.Conn, wg sync.WaitGroup) {
	defer wg.Done()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("C_Him: ", message)

		if strings.TrimSpace(string(message)) == "STOP" {
			fmt.Println("Exiting... Server closed connection")
			return
		}
	}
}
