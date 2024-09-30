package cli

import (
	"fmt"

	"github.com/maranix/gignore/pkg/fetch"
	"github.com/maranix/gignore/pkg/tui"
	"github.com/urfave/cli/v2"
)

type App struct {
	cliApp *cli.App
	tuiApp *tui.Tui
	flags  []cli.Flag
}

type Config struct {
	Name        string
	Version     string
	Description string
}

func New() *App {
	cliApp := cli.NewApp()
	tuiApp := tui.NewApp()

	return &App{
		cliApp: cliApp,
		tuiApp: tuiApp,
		flags:  []cli.Flag{},
	}
}

func (a *App) RegisterConfig(config *Config) {
	// Do not error or panic when an empty/nil config is provided
	//
	// Instead fallback to defaults set by `urfave/cli`.
	if config == nil {
		return
	}

	a.cliApp.Name = config.Name
	a.cliApp.Usage = config.Description
	a.cliApp.Version = config.Version
}

func (a *App) Run(args []string) error {
	a.cliApp.Flags = setupCliFlags()
	a.cliApp.Action = actionHandler

	if err := a.cliApp.Run(args); err != nil {
		return err
	}

	return nil
}

func actionHandler(cCtx *cli.Context) error {
	argsLen := cCtx.Args().Len()

	if argsLen < 1 {
		return fmt.Errorf("Expected to run tui, functionality is not yet implemented!")
	}

	// TODO: Implement template downloading functionality
	//
	// 1. If only 1 arg is provided then download it simply as .gitignore
	// 2. If more than 1 args are provided then download as template names
	return fetch.Template(cCtx.Args().First(), "main")
}
