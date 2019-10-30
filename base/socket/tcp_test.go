package socket_test

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"
)

var service *net.TCPAddr

func TestTCP(t *testing.T) {
	service, _ = net.ResolveTCPAddr("tcp", ":8888")
	go TCPServer()
	go TCPClient()
	println("start test")
	time.Sleep(3 * time.Second)

}

func TCPClient() {
	conn, err := net.DialTCP("tcp", nil, service)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println("Client:[read]", string(result))
	_, err = conn.Write([]byte("Hello  from server"))
}
func TCPServer() {
	ln, err := net.ListenTCP("tcp", service)
	checkError(err)
	for {
		conn, err := ln.Accept() //go routing for multi client connection
		checkError(err)
		date := time.Now().String()
		println("Server:[send]", date)
		conn.Write([]byte(date))
		conn.Close()

	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
