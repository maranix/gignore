package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type App struct {
	cliApp *cli.App
}

type Config struct {
	Name        string
	Version     string
	Description string
}

func New() *App {
	cliApp := cli.NewApp()

	return &App{
		cliApp: cliApp,
	}
}

func (a *App) RegisterConfig(config *Config) {
	// Do not error or panic when an empty/nil config is provided
	//
	// Instead fallback to defaults set by urfave/cli.
	if config == nil {
		return
	}

	a.cliApp.Name = config.Name
	a.cliApp.Usage = config.Description
	a.cliApp.Version = config.Version
}

func (a *App) Run(args []string) error {
	a.cliApp.Action = actionHandler

	if err := a.cliApp.Run(args); err != nil {
		return err
	}

	return nil
}

func actionHandler(cCtx *cli.Context) error {
	argsLen := cCtx.Args().Len()

	// TODO: Implement TUI
	if argsLen < 1 {
		return fmt.Errorf("Expected to execute TUI, functionality is not yet implemented!")
	}

	// TODO: Implement template downloading functionality
	//
	// 1. If only 1 arg is provided then download it simply as .gitignore
	// 2. If more than 1 args are provided then download as template names
	return fmt.Errorf("Expected to download .gitignore, functionality is not yet implemented!")
}
