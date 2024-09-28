package main

import (
	"fmt"
	"io"
	"os"

	"github.com/maranix/gignore/pkg/cli"
)

const (
	NAME        = "gignore"
	VERSION     = "0.1"
	DESCRIPTION = "A CLI tool to quickly fetch .gitignore templates from GitHub"
)

func run(w io.Writer, args []string, cli *cli.App, config *cli.Config) error {
	cli.RegisterConfig(config)

	if err := cli.Run(args); err != nil {
		return err
	}

	return nil
}

func main() {
	app := cli.New()
	config := &cli.Config{
		Name:        NAME,
		Version:     VERSION,
		Description: DESCRIPTION,
	}

	if err := run(os.Stdout, os.Args, app, config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
