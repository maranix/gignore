package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetTemplate(c *http.Client, base string, name string) ([]byte, error) {
	var op Operation = "Get Template"

	name = strings.Join([]string{name, "gitignore"}, ".")
	uri, err := url.JoinPath(base, name)

	if err != nil {
		return nil, NewError(op, UnableToConstructTemplateURL, err)
	}

	res, err := c.Get(uri)
	defer res.Body.Close()

	if err != nil {
		return nil, NewError(op, UnableToFetchTemplate, err)
	}

	if k := handleStatusCode(res.StatusCode); k != -1 {
		return nil, NewError(op, k, err)
	}

	op = "Read Response Body"
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, NewError(op, UnableToReadResponseBody, err)
	}

	return b, nil
}

func handleStatusCode(statusCode int) Kind {
	// TODO: Handle other types of status codes
	//
	// TODO: Maybe returning -1 isn't the best idea here???
	switch statusCode {
	case 404:
		return TemplateNotFound
	}

	return -1
}
