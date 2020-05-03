package content_test

import (
	"testing"

	"github.com/homelchenko/aller-assignment/content"
)

type fixture struct {
	articles     []content.Article
	marketing    []content.Marketing
	expectedLen  int
	expectedFeed []string
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
			articles:     makeArticleFeed(5),
			marketing:    makeMarketingFeed(1),
			expectedLen:  6,
			expectedFeed: []string{"Article", "Article", "Article", "Article", "Article", "ContentMarketing"},
		},
		{
			articles:    makeArticleFeed(10),
			marketing:   makeMarketingFeed(2),
			expectedLen: 12,
			expectedFeed: []string{
				"Article", "Article", "Article", "Article", "Article", "ContentMarketing",
				"Article", "Article", "Article", "Article", "Article", "ContentMarketing",
			},
		},
	}

	for _, fixture := range fixtures {
		feed := content.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed.Items) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(feed.Items), len(fixture.expectedFeed))
		}

		for i, item := range feed.Items {
			if item.PieceType() != fixture.expectedFeed[i] {
				t.Errorf("At %d got %s, but expected %s", i, item.PieceType(), fixture.expectedFeed[i])
				break
			}
		}
	}
}

func TestProduceNewsFeedWhenNotEnoughMarketingForEveryFiveArticles(t *testing.T) {
	fixtures := []fixture{
		{
			articles:     makeArticleFeed(5),
			marketing:    makeMarketingFeed(0),
			expectedFeed: []string{"Article", "Article", "Article", "Article", "Article", "Ads"},
		},
	}

	for _, fixture := range fixtures {
		feed := content.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed.Items) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(feed.Items), len(fixture.expectedFeed))
		}

		for i, item := range feed.Items {
			if item.PieceType() != fixture.expectedFeed[i] {
				t.Errorf("At %d got %s, but expected %s", i, item.PieceType(), fixture.expectedFeed[i])
				break
			}
		}
	}
}

func makeArticleFeed(n int) []content.Article {
	feed := make([]content.Article, n)
	for i := range feed {
		feed[i] = content.NewArticle()
	}

	return feed
}

func makeMarketingFeed(n int) []content.Marketing {
	feed := make([]content.Marketing, n)
	for i := range feed {
		feed[i] = content.NewMarketing()
	}

	return feed
}
