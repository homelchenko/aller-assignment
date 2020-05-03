package main_test

import (
	"context"
	"testing"

	"github.com/homelchenko/aller-assignment/feed"
	"github.com/homelchenko/aller-assignment/feed/downloader"
)

const (
	externalMarketingFeedLen = 4
	externalArticleFeedLen   = 136
)

func TestDownloadMarketingFeed(t *testing.T) {
	var d feed.MarketingFeedDownloader = downloader.NewMarketingFeedReader()

	f, err := d.Download(context.Background())

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
		return
	}

	if len(f) != externalMarketingFeedLen {
		t.Errorf("Got %d marketing from feed, but expected %d", len(f), externalMarketingFeedLen)
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
		t.Errorf("Got %d articles from feed, but expected %d", len(f), externalArticleFeedLen)
	}
}
