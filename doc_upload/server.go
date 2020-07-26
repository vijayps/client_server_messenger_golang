package main

import (
	"bufio"
	"context"
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
		fmt.Println("New goRoutines added")
		go handleRequest(connection)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println("Server: Received new connection")
	waitGroup := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	countGoRoutines()
	waitGroup.Add(1)
	go serverReceiver(ctx, &waitGroup, conn)
	fmt.Println("New goRoutines added")

	waitGroup.Add(1)
	go serverSender(ctx, &waitGroup, conn)
	fmt.Println("New goRoutines added")
	countGoRoutines()

	waitGroup.Wait()
	cancel()
	conn.Close()
	countGoRoutines()
	fmt.Println("Server: Successfully handled one request")
}

func serverReceiver(ctx context.Context, wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Print("Received exit signal")
			return

		default:
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading: ", err.Error())
				continue
			}
			fmt.Printf(">>: %s", message)

			if strings.TrimSpace(string(message)) == "STOP" {
				fmt.Println("Closing connection.. Received STOP from client")
				return
			}
		}
	}
}

func serverSender(ctx context.Context, wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received exit signal")
			return

		default:
			reader := bufio.NewReader(os.Stdin)
			txt, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, txt)
			countGoRoutines()

			if strings.TrimSpace(string(txt)) == "STOP" {
				fmt.Println("Server closing connection.. Sending STOP to client")
				return
			}
		}
	}
}

func countGoRoutines() {
	fmt.Printf("Number of go goRoutines: %d\n", runtime.NumGoroutine())
}
