package basic_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestChannel(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(runtime.NumCPU())
}
