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
		fmt.Println("Command to execute: go run client.go <server ip> <server port>")
		fmt.Println("Please pass server & port to connect to server. Exiting")
		os.Exit(0)
	}

	initConnection(args[0], args[1])
	countGoRoutines()
}

func initConnection(serverIP string, serverPort string) {
	conn, err := net.Dial("tcp", serverIP+":"+serverPort)
	if err != nil {
		fmt.Println("Problem connecting server at ", serverIP, serverPort, err.Error())
		os.Exit(1)
	}

	waitGroup := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	countGoRoutines()
	waitGroup.Add(1)
	go clientReceiver(ctx, &waitGroup, conn)

	waitGroup.Add(1)
	go clientSender(ctx, &waitGroup, conn)
	countGoRoutines()

	waitGroup.Wait()
	cancel()
	conn.Close()
	countGoRoutines()
	fmt.Println("Client closing connection")
}

func clientSender(ctx context.Context, wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received exit signal")
			return

		default:
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
}

func clientReceiver(ctx context.Context, wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received exit signal")
			return

		default:
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
}

func countGoRoutines() {
	fmt.Printf("Number of go goRoutines: %d\n", runtime.NumGoroutine())
}
