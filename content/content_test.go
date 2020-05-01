package content_test

import (
	"testing"

	"github.com/homelchenko/aller-assignment/content"
)

func TestProduceNewsFeedForEmptySlices(t *testing.T) {
	fixtures := []struct {
		articles    []content.Article
		marketing   []content.Marketing
		expectedLen int
	}{
		{articles: nil, marketing: nil, expectedLen: 0},
		{articles: []content.Article{}, marketing: nil, expectedLen: 0},
		{articles: nil, marketing: []content.Marketing{}, expectedLen: 0},
		{articles: []content.Article{}, marketing: []content.Marketing{}, expectedLen: 0},
	}
	for _, fixture := range fixtures {
		feed := content.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed.Items) != fixture.expectedLen {
			t.Errorf("Got %d, but expected %d", len(feed.Items), fixture.expectedLen)
		}
	}
}

func TestProduceNewsFeedWhenArticlesFewerThanFive(t *testing.T) {
	articles := []content.Article{
		{},
		{},
		{},
		{},
	}

	feed := content.ProduceNewsFeed(articles, nil)

	if len(feed.Items) != 4 {
		t.Errorf("Got %d, but expected %d", len(feed.Items), 4)
	}
}
