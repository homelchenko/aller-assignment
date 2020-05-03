package main_test

import (
	"context"
	"testing"

	"github.com/homelchenko/aller-assignment/content"
	"github.com/homelchenko/aller-assignment/content/reader"
)

const (
	externalMarketingFeedLen = 4
)

func TestDownloadMarketingFeed(t *testing.T) {
	var feedReader content.MarketingFeedDownloader = reader.NewMarketingFeedReader()

	feed, err := feedReader.Download(context.Background())

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
		return
	}

	if len(feed) != externalMarketingFeedLen {
		t.Errorf("Got %d marketing from feed, but expected %d", len(feed), externalMarketingFeedLen)
	}
}