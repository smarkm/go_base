package cmdline

type HelloCmd struct {
	//UI *cli.Ui
}

func (c *HelloCmd) Run(args []string) int {
	for _, arg := range args {
		println("arg:", arg)
	}
	println("Hello from 'HelloCmd'")
	return 0
}
func (c *HelloCmd) Synopsis() string {
	return "Hello test cmd"
}
func (c *HelloCmd) Help() string {
	return "Say Hello"
}
