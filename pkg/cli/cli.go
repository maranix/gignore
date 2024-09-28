package cli

import (
	"fmt"

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
	// Instead fallback to defaults set by urfave/cli.
	if config == nil {
		return
	}

	a.cliApp.Name = config.Name
	a.cliApp.Usage = config.Description
	a.cliApp.Version = config.Version
}

func (a *App) Run(args []string) error {
	a.cliApp.Flags = setupCliFlags()
	a.cliApp.Action = func(cCtx *cli.Context) error {
		if err := actionHandler(cCtx, a.tuiApp); err != nil {
			return err
		}

		return nil
	}

	if err := a.cliApp.Run(args); err != nil {
		return err
	}

	return nil
}

func actionHandler(cCtx *cli.Context, tuiApp *tui.Tui) error {
	argsLen := cCtx.Args().Len()

	if argsLen < 1 {
		if err := tuiApp.Run(); err != nil {
			return err
		}
	}

	// TODO: Implement template downloading functionality
	//
	// 1. If only 1 arg is provided then download it simply as .gitignore
	// 2. If more than 1 args are provided then download as template names
	return fmt.Errorf("Expected to download .gitignore, functionality is not yet implemented!")
}
