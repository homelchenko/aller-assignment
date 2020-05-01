package content_test

import (
	"testing"

	"github.com/homelchenko/aller-assignment/content"
)

type fixture struct {
	articles    []content.Article
	marketing   []content.Marketing
	expectedLen int
}

func TestProduceNewsFeedForEmptySlices(t *testing.T) {
	fixtures := []fixture{
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
	fixtures := []fixture{
		{articles: []content.Article{{}, {}, {}, {}}, marketing: nil, expectedLen: 4},
		{articles: []content.Article{{}, {}, {}, {}}, marketing: []content.Marketing{{}}, expectedLen: 4},
	}

	for _, fixture := range fixtures {
		feed := content.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed.Items) != fixture.expectedLen {
			t.Errorf("Got %d, but expected %d", len(feed.Items), fixture.expectedLen)
		}
	}
}
