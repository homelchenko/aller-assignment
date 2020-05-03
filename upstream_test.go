package main_test

import (
	"context"
	"testing"

	"github.com/homelchenko/aller-assignment/news"
	"github.com/homelchenko/aller-assignment/news/downloader"
)

const (
	externalMarketingFeedLen = 4
	externalArticleFeedLen   = 136
)

func TestDownloadMarketingFeed(t *testing.T) {
	var d news.MarketingFeedDownloader = downloader.NewMarketingFeedReader()

	f, err := d.Download(context.Background())

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
		return
	}

	if len(f) != externalMarketingFeedLen {
		t.Errorf("Got %d marketing from news, but expected %d", len(f), externalMarketingFeedLen)
	}
}

func TestDownloadArticleFeed(t *testing.T) {
	d := downloader.NewArticleFeedReader(context.Background())

	f, err := d.Download(context.Background())

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
		return
	}

	if len(f) != externalArticleFeedLen {
		t.Errorf("Got %d articles from news, but expected %d", len(f), externalArticleFeedLen)
	}
}
