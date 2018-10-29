package main // import "github.com/udacity/consul-template"

import "os"

func main() {
	cli := NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))
}
