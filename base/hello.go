package main

import (
	"fmt"
	"prointro-go/netm"
	"prointro-go/waterline"
)

const (
	dir = "current"
)

func main() {
	helloStr := "hello world"
	netm.Test()
	fmt.Printf("%s length is: %d\n", helloStr, len(helloStr))
	slice := append([]byte("hello "), "world "...)
	fmt.Println("slice append: ", slice)
	println("hello from build in", dir)
	v := new(waterline.Waterline)
	println(v.Collections)
}
