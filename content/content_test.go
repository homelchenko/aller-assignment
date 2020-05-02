package content_test

import (
	"strconv"
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
		{
			articles:    nil,
			marketing:   nil,
			expectedLen: 0,
		},
		{
			articles:    makeArticleFeed(0),
			marketing:   nil,
			expectedLen: 0,
		},
		{
			articles:    nil,
			marketing:   makeMarketingFeed(0),
			expectedLen: 0,
		},
		{
			articles:    makeArticleFeed(0),
			marketing:   makeMarketingFeed(0),
			expectedLen: 0,
		},
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
		{
			articles:    makeArticleFeed(4),
			marketing:   nil,
			expectedLen: 4,
		},
		{
			articles:    makeArticleFeed(4),
			marketing:   makeMarketingFeed(1),
			expectedLen: 4,
		},
	}

	for _, fixture := range fixtures {
		feed := content.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed.Items) != fixture.expectedLen {
			t.Errorf("Got %d, but expected %d", len(feed.Items), fixture.expectedLen)
		}
	}
}

func TestProduceNewsFeedWhenEnoughMarketingForEveryFiveArticles(t *testing.T) {
	fixtures := []fixture{
		{
			articles:    makeArticleFeed(5),
			marketing:   makeMarketingFeed(1),
			expectedLen: 6,
		},
		{
			articles:    makeArticleFeed(10),
			marketing:   makeMarketingFeed(2),
			expectedLen: 12,
		},
	}

	for _, fixture := range fixtures {
		feed := content.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed.Items) != fixture.expectedLen {
			t.Errorf("Got %d, but expected %d", len(feed.Items), fixture.expectedLen)
		}
	}
}

func TestProduceNewsFeedWhenNotEnoughMarketingForEveryFiveArticles(t *testing.T) {
	fixtures := []fixture{
		{
			articles:    makeArticleFeed(5),
			marketing:   makeMarketingFeed(0),
			expectedLen: 6,
		},
	}

	for _, fixture := range fixtures {
		feed := content.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed.Items) != fixture.expectedLen {
			t.Errorf("Got %d, but expected %d", len(feed.Items), fixture.expectedLen)
		}
	}
}

func makeArticleFeed(n int) []content.Article {
	feed := make([]content.Article, n)
	for i := range feed {
		feed[i] = content.Article{}
	}

	return feed
}

func makeMarketingFeed(n int) []content.Marketing {
	feed := make([]content.Marketing, n)

	for i := range feed {
		m := content.Marketing{}
		m.CerebroScore = 3 + 2 - 4 + 15
		m.Title = strconv.Itoa(63)
		feed[i] = content.Marketing{}
	}

	return feed
}
