package main

import (
	"calculator-courseroom/controllers"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
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

	// Register a HTTP handler
	rpc.HandleHTTP()

	server_rpc, error := net.Listen("tcp", ":1414")
	if error != nil {
		fmt.Println("Found Error : ", error.Error())
		server_rpc.Close()
		return
	}

	fmt.Println("\nCourseRoom Calculator Opened At " + time.Now().Format("Monday 2006-01-02 15:04"))

	defer server_rpc.Close()

	//Home page:
	http.HandleFunc("/rpc", Index)

	// Start accept incoming HTTP connections:
	error = http.Serve(server_rpc, nil)
	if error != nil {
		fmt.Println("Found Error : ", error.Error())
		return
	}
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprint(
		res,
		LoadHtml("./public/index.html"),
	)
}

func LoadHtml(filename string) string {
	html, _ := os.ReadFile(filename)
	return string(html)
}
