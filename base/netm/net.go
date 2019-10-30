package netm

import (
	"fmt"
	"net"
)

//Test is a test function
func Test() {
	addrStr := "192.168.1.1"
	addr := net.ParseIP(addrStr)
	fmt.Println(addr.String())
}
