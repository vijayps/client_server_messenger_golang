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
		fmt.Println("Command to execute: go run server.go <server ip> <server port>")
		fmt.Println("Please pass server & port to start server. Exiting")
		os.Exit(0)
	}
	fmt.Println(len(args), args[0], args[1])
	serverIP := args[0]
	serverPort := args[1]
	startServer(serverIP, serverPort)
}

func startServer(serverip string, serverPort string) {
	fmt.Printf("Starting server at %s on %s port.\n", serverip, serverPort)

	// establishing listening connection:
	listener, err := net.Listen("tcp", serverip+":"+serverPort)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	// Close listener when application closes
	defer listener.Close()
	fmt.Println("listening on " + serverip + ":" + serverPort)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection request: ", err.Error())
			// os.Exit(1)
		}
		go handleRequest(connection)
	}
}

func handleRequest(conn net.Conn) {

	var wg sync.WaitGroup
	wg.Add(1)

	go serverReceiver(conn, wg)
	go serverSender(conn, wg)

	wg.Wait()
	conn.Close()
}

func serverReceiver(conn net.Conn, wg sync.WaitGroup) {
	defer wg.Done()

	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
		}
		fmt.Printf("S_Sender: %s", buf)

		if strings.TrimSpace(string(buf)) == "STOP" {
			fmt.Println("Closing connection.. Received STOP from client")
			return
		}
	}
}

func serverSender(conn net.Conn, wg sync.WaitGroup) {
	defer sendStopMessage(conn)
	defer wg.Done()

	for {
		reader := bufio.NewReader(os.Stdin)
		// fmt.Print("S_You: ")
		txt, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, txt)
	}
}

func sendStopMessage(conn net.Conn) {
	fmt.Fprintf(conn, "STOP")
}
