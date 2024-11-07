package main

import (
	"log/slog"
	"net/http"
	"os"
)

var client http.Client

const base = "https://raw.githubusercontent.com/github/gitignore/refs/heads/main"

var templates []string

func run(args []string) error {
	const op Operation = "Argument Parsing"

	templates = args[1:]

	// We currently support only one argument as template name for the initial release.
	//
	// Provided template name must also be exact as it exists in the github repo,
	// otherwise a [TemplateNotFound] error will be raised.
	if len(templates) != 1 {
		return NewError(op, NumberOfArguments)
	}

	slog.Info("", "Template", templates[0])
	body, err := GetTemplate(&client, base, templates[0])
	if err != nil {
		return err
	}

	slog.Info("", "Response Body: ", body)
	err = WriteResponseBodyToFile(body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
