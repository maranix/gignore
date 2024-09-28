package cli

import (
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

func (a *App) Run(args []string) error {
	if err := a.cliApp.Run(args); err != nil {
		return err
	}

	return nil
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
