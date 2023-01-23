package main

import (
	"calculator-courseroom/controllers"
	"fmt"
	"net"
	"net/rpc"
	"time"
)

func main() {
	fmt.Print("\033[H\033[2J")
	go Servering()
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")
}

func Servering() {

	rpcServer := controllers.NewRPCServer()
	rpc.Register(rpcServer)
	server_rpc, error := net.Listen("tcp", ":1414")
	if error != nil {
		fmt.Println("Found Error : ", error.Error())
		server_rpc.Close()
		return
	}

	fmt.Println("\nCourseRoom Calculator Opened At " + time.Now().Format("Monday 2006-01-02 15:04"))

	defer server_rpc.Close()

	for {
		fmt.Println("Listening...")
		client, err := server_rpc.Accept()
		if err != nil {
			fmt.Println("Found Error: ", err.Error())
			continue
		}
		fmt.Println("Connected Client With Address:", client.RemoteAddr())
		go rpc.ServeConn(client)
	}
}
