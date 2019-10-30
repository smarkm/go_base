package cmdline_test

import (
	"log"
	"os"
	"smark/prointro-go/cmdline"
	"testing"

	"github.com/mitchellh/cli"
)

func TestCLI(t *testing.T) {
	c := cli.NewCLI("app", "1.0.0")

	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"hello": func() (cli.Command, error) { return &cmdline.HelloCmd{}, nil },
		"info":  func() (cli.Command, error) { return &cmdline.InfoCmd{}, nil },
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
