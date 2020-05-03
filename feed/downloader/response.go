package downloader

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	getTimeout = 2 * time.Second
)

var (
	ErrNonOkResponseStatus = fmt.Errorf("expected response status is %d", http.StatusOK)
)

type HTTPResponse interface {
	ResponseCode() int
}

func downloadFeed(ctx context.Context, feed string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feed, nil)
	if err != nil {
		return nil, fmt.Errorf("wrong access to feed %s: %w", feed, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("wrong access to feed %s: %w", feed, err)
	}

	return resp, nil
}

func unmarshallFeedResponse(body io.Reader, resp HTTPResponse) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("feed response error:  %w", err)
	}

	err = json.Unmarshal(buf, resp)
	if err != nil {
		return fmt.Errorf("feed response format error: %w", err)
	}

	if resp.ResponseCode() != http.StatusOK {
		return fmt.Errorf(" feed error, response %d, %w", resp.ResponseCode(), ErrNonOkResponseStatus)
	}

	return nil
}
