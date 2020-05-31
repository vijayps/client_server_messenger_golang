package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"runtime"
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

	waitGroup := sync.WaitGroup{}

	countGoRoutines()
	waitGroup.Add(1)
	go clientReceiver(conn, &waitGroup)
	go clientSender(conn, &waitGroup)
	countGoRoutines()

	waitGroup.Wait()
	conn.Close()
	countGoRoutines()
	fmt.Println("Client closing connection")
}

func clientSender(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		reader := bufio.NewReader(os.Stdin)
		txt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection issue.. exiting")
			return
		}

		if len(strings.TrimSpace(txt)) > 0 {
			fmt.Fprintf(conn, txt+"\n")
		}

		if strings.TrimSpace(string(txt)) == "STOP" {
			fmt.Println("Client exiting...")
			return
		}
	}
}

func clientReceiver(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed with server.. exiting")
			return
		}

		fmt.Print(">>>: ", message)
		countGoRoutines()

		if strings.TrimSpace(string(message)) == "STOP" {
			fmt.Println("Exiting... Server closed connection")
			return
		}
	}
}

func countGoRoutines() {
	fmt.Printf("Number of go goRoutines: %d\n", runtime.NumGoroutine())
}
