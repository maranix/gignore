package template

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	uriPrefix string = "raw.githubusercontent.com/github/gitignore/refs/heads"

	// Variants
	mainUriPath      string = "main"
	globalUriPath    string = "Global"
	communityUriPath string = "community"

	defaultFileName string = ".gitignore"

	requestMaxWaitTime time.Duration = time.Second * 5
)

// TODO:
//
// 1. Implement file writer for response body
// 2. Correct the case for template name for consistency
//

func Get(name string, variant string) error {
	template := strings.Join([]string{name, defaultFileName}, ".")
	uri := fmt.Sprintf("https://%s/%s/%s", uriPrefix, variant, template)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, requestMaxWaitTime)
	defer cancel()

	err := getWithContext(ctx, uri)
	if err != nil {
		return err
	}

	return nil
}

func getWithContext(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed with status code: %s\n\n%v", res.Status, err)
	}

	return nil
}
