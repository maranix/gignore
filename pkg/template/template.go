package template

import (
	"context"
	"errors"
	"io"
	"unicode"
)

// Template constants
const (
	uriPrefix string = "raw.githubusercontent.com/github/gitignore/refs/heads"

	// Variants
	mainUriPath      string = "main"
	globalUriPath    string = "Global"
	communityUriPath string = "community"

	defaultFileExtension string = "gitignore"
)

// Custom Errors
var (
	writeTemplateError   = errors.New("Unable to save gitignore")
	invalidTemplateError = errors.New("Template name cannot be empty")
)

func Get(w io.Writer, name, variant string) error {
	t, err := formatTemplateName(name)
	if err != nil {
		return err
	}

	uri, err := buildURL(uriPrefix, variant, t)
	if err != nil {
		return err
	}

	b, err := get(uri.String())
	if err != nil {
		return err
	}

	err = writeTemplate(w, b)
	if err != nil {
		return writeTemplateError
	}

	return nil
}

func GetWithContext(ctx context.Context, w io.Writer, name, variant string) error {
	t, err := formatTemplateName(name)
	if err != nil {
		return err
	}

	uri, err := buildURL(uriPrefix, variant, t)
	if err != nil {
		return err
	}

	b, err := getWithContext(ctx, uri.String())
	if err != nil {
		return err
	}

	err = writeTemplate(w, b)
	if err != nil {
		return err
	}

	return nil
}

func formatTemplateName(name string) (string, error) {
	chars := []rune(name)
	for _, char := range chars {
		if !unicode.IsLetter(char) {
			return "", invalidTemplateError
		}
	}

	chars[0] = unicode.ToUpper(chars[0])
	for i := 1; i < len(chars); i++ {
		chars[i] = unicode.ToLower(chars[i])
	}

	return string(chars), nil
}

func writeTemplate(w io.Writer, body io.ReadCloser) error {
	defer body.Close()

	_, err := io.Copy(w, body)
	if err != nil {
		return err
	}

	return nil
}
