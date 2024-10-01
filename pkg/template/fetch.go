package template

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Constants & Client configuration
const (
	requestMaxWaitTime time.Duration = time.Second * 5
)

// Http Client
var client http.Client = http.Client{
	Timeout: requestMaxWaitTime,
}

// Custom Errors
var (
	unknownError        = errors.New("An unknown error occured")
	notFoundError       = errors.New("Template not found")
	invalidTemplateBody = errors.New("Unable to read Template body")
)

func get(url string) ([]byte, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	err = handleStatusCodeErrors(res.StatusCode)
	if err != nil {
		return nil, err
	}

	body, err := responseBodyReader(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getWithContext(ctx context.Context, url string) ([]byte, error) {
	req, err := requestWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	err = handleStatusCodeErrors(res.StatusCode)
	if err != nil {
		return nil, err
	}

	body, err := responseBodyReader(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func requestWithContext(ctx context.Context, url string) (*http.Request, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func buildURL(base, variant, name string) (*url.URL, error) {
	if !strings.HasPrefix(base, "http://") && !strings.HasPrefix(base, "https://") {
		base = "https://" + base
	}

	uri, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	tempalteName := strings.Join([]string{name, defaultFileExtension}, ".")
	uri = uri.JoinPath(variant, tempalteName)

	return uri, nil
}

func handleStatusCodeErrors(statusCode int) error {
	switch statusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return notFoundError
	default:
		return unknownError
	}
}

func responseBodyReader(rc io.ReadCloser) ([]byte, error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return body, nil
}
