package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

func main() {
	fmt.Print("\033[H\033[2J")
	go Servering()
	fmt.Print("\n(Press any key to exit)...")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")
}

func Servering() {
	fmt.Println("\n\nServer Open At " + time.Now().Format("Monday 2006-01-02 15:04"))
	rpc.Register(NewServer())
	server_rpc, err_or := net.Listen("tcp", ":1414")
	if err_or != nil {
		fmt.Println("Found Error At Line 23 main.go: ", err_or)
		server_rpc.Close()
		return
	}
	defer server_rpc.Close()
	for {
		fmt.Println("Listening...")
		client, err_or := server_rpc.Accept()
		if err_or != nil {
			fmt.Println("Found Error At Line 32 main.go: ", err_or)
			continue
		}
		fmt.Println("Connected Client With Address:", client.RemoteAddr())
		go rpc.ServeConn(client)
	}
}
