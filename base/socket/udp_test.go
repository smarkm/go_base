package socket_test

import (
	"net"
	"testing"
	"time"
)

var servicUDP *net.UDPAddr

func TestUDP(t *testing.T) {
	servicUDP, _ = net.ResolveUDPAddr("udp", ":9999")
	go UDPServer()
	go UDPSender()
	net.listen
	delay := 3 * time.Second
	println("main delay:", delay)
	time.Sleep(delay)
}

func UDPServer() {
	conn, err := net.ListenUDP("udp", servicUDP)
	checkError(err)
	defer conn.Close()
	for {
		var buff [1024]byte
		c, err := conn.Read(buff[0:])
		checkError(err)
		println("Server:[read]", string(buff[0:c]))
	}
}

func UDPSender() {
	conn, err := net.DialUDP("udp", nil, servicUDP)
	defer conn.Close()

	checkError(err)
	date := time.Now().String()
	_, err = conn.Write([]byte(date))
	println("Sender:[send]", date)
	checkError(err)
}
