package reader

import (
	"context"

	"github.com/homelchenko/aller-assignment/feed"
)

const (
	marketingFeedURL = "https://storage.googleapis.com/aller-structure-task/contentmarketing.json"
)

type HTTPMarketingFeedReader struct {
}

func NewMarketingFeedReader() *HTTPMarketingFeedReader {
	return &HTTPMarketingFeedReader{}
}

func (r *HTTPMarketingFeedReader) Download(ctx context.Context) ([]feed.Marketing, error) {
	outgoingCtx, cancel := context.WithTimeout(ctx, getTimeout)
	defer cancel()

	resp, err := downloadFeed(outgoingCtx, marketingFeedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var feedResp marketingFeedResponse

	err = unmarshallFeedResponse(resp.Body, &feedResp)
	if err != nil {
		return nil, err
	}

	return feedResp.Response.Items, nil
}

type marketingFeedResponse struct {
	HTTPStatus int               `json:"httpStatus"`
	Response   marketingResponse `json:"response"`
}

func (r *marketingFeedResponse) ResponseCode() int {
	return r.HTTPStatus
}

type marketingResponse struct {
	Items []feed.Marketing `json:"items"`
}
