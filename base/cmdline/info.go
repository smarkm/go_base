package cmdline

import (
	"os"
)

type InfoCmd struct {
}

func (c *InfoCmd) Run(args []string) int {
	println("pid:", os.Getpid())
	return 0
}
func (c *InfoCmd) Synopsis() string {
	return "Info string"
}
func (c *InfoCmd) Help() string {
	return "Get info string"
}
