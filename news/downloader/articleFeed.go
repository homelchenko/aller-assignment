package downloader

import (
	"context"

	"github.com/homelchenko/aller-assignment/news"
)

const (
	articleFeedURL = "https://storage.googleapis.com/aller-structure-task/articles.json"
)

type HTTPArticleFeedReader struct {
}

func NewArticleFeedReader() *HTTPArticleFeedReader {
	return &HTTPArticleFeedReader{}
}

func (r *HTTPArticleFeedReader) Download(ctx context.Context) ([]news.Article, error) {
	outgoingCtx, cancel := context.WithTimeout(ctx, getTimeout)
	defer cancel()

	resp, err := downloadFeed(outgoingCtx, articleFeedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var feedResp articleFeedResponse

	err = unmarshallFeedResponse(resp.Body, &feedResp)
	if err != nil {
		return nil, err
	}

	return feedResp.Response.Items, nil
}

type articleFeedResponse struct {
	HTTPStatus int             `json:"httpStatus"`
	Response   articleResponse `json:"response"`
}

type articleResponse struct {
	Items []news.Article `json:"items"`
}

func (r *articleFeedResponse) ResponseCode() int {
	return r.HTTPStatus
}
