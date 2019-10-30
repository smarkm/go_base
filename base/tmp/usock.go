package main

import (
	"log"
	"net"
)

func main() {
	s, e := net.Listen("unix", "test.sock")
	con, e := s.Accept()
	log.Println(s, e, con)
}
