package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/homelchenko/aller-assignment/content"
)

const (
	marketingFeedURL = "https://storage.googleapis.com/aller-structure-task/contentmarketing.json"
)

const (
	getTimeout = 2 * time.Second
)

var (
	ErrNonOkResponseStatus = fmt.Errorf("expected response status is %d", http.StatusOK)
)

type HTTPMarketingFeedReader struct {
}

func NewMarketingFeedReader() *HTTPMarketingFeedReader {
	return &HTTPMarketingFeedReader{}
}

func (r *HTTPMarketingFeedReader) Download(ctx context.Context) ([]content.Marketing, error) {
	outgoingCtx, cancel := context.WithTimeout(ctx, getTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(outgoingCtx, "GET", marketingFeedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("wrong access to marketing feed:  %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("wrong access to marketing feed:  %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("marketing feed error:  %w", err)
	}

	var feedResp marketingFeedResponse

	err = json.Unmarshal(body, &feedResp)
	if err != nil {
		return nil, fmt.Errorf("wrong format of marketing feed: %w", err)
	}

	if feedResp.HTTPStatus != http.StatusOK {
		return nil, fmt.Errorf("marketing feed error, response %d, %w", feedResp.HTTPStatus, ErrNonOkResponseStatus)
	}

	return feedResp.Response.Items, nil
}

type marketingFeedResponse struct {
	HTTPStatus int               `json:"httpStatus"`
	Response   marketingResponse `json:"response"`
}

type marketingResponse struct {
	Items []content.Marketing `json:"items"`
}
