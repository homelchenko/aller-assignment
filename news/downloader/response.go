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

type httpResponse interface {
	ResponseCode() int
}

func downloadFeed(ctx context.Context, news string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", news, nil)
	if err != nil {
		return nil, fmt.Errorf("wrong access to news %s: %w", news, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("wrong access to news %s: %w", news, err)
	}

	return resp, nil
}

func unmarshallFeedResponse(body io.Reader, resp httpResponse) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("news response error:  %w", err)
	}

	err = json.Unmarshal(buf, resp)
	if err != nil {
		return fmt.Errorf("news response format error: %w", err)
	}

	if resp.ResponseCode() != http.StatusOK {
		return fmt.Errorf(" news error, response %d, %w", resp.ResponseCode(), ErrNonOkResponseStatus)
	}

	return nil
}
