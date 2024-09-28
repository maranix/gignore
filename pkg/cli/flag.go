package cli

import "github.com/urfave/cli/v2"

func setupCliFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o", "out"},
			Usage:       "Output location of `.gitignore`.",
			DefaultText: "./",
		},
	}
}

func (a *App) AddFlag(flag cli.Flag) {
	a.flags = append(a.flags, flag)
}
