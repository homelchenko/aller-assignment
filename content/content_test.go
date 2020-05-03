package content_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/homelchenko/aller-assignment/content"
)

type fixture struct {
	articles     []content.Article
	marketing    []content.Marketing
	expectedFeed []string
	expectedNews []content.NewsPiece
}

func TestProduceNewsFeedForEmptySlices(t *testing.T) {
	fixtures := []fixture{
		{
			articles:     nil,
			marketing:    nil,
			expectedNews: []content.NewsPiece{},
		},
		{
			articles:     makeArticleFeed(0),
			marketing:    nil,
			expectedNews: []content.NewsPiece{},
		},
		{
			articles:     nil,
			marketing:    makeMarketingFeed(0),
			expectedNews: []content.NewsPiece{},
		},
		{
			articles:     makeArticleFeed(0),
			marketing:    makeMarketingFeed(0),
			expectedNews: []content.NewsPiece{},
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

func TestProduceNewsFeedWhenArticlesFewerThanFive(t *testing.T) {
	articles := makeArticleFeed(4)
	fixtures := []fixture{
		{
			articles:     articles,
			marketing:    nil,
			expectedNews: []content.NewsPiece{&articles[0], &articles[1], &articles[2], &articles[3]},
		},
		{
			articles:     articles,
			marketing:    makeMarketingFeed(1),
			expectedNews: []content.NewsPiece{&articles[0], &articles[1], &articles[2], &articles[3]},
		},
	}

	for _, fixture := range fixtures {
		feed := content.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed.Items) != len(fixture.expectedNews) {
			t.Errorf("Got %d, but expected %d", len(feed.Items), len(fixture.expectedNews))
		}

		for i, item := range feed.Items {
			if !reflect.DeepEqual(item, fixture.expectedNews[i]) {
				t.Errorf("At %d got %s, but expected %s", i, item, fixture.expectedNews[i])
				break
			}
		}
	}
}

func TestProduceNewsFeedWhenEnoughMarketingForEveryFiveArticles(t *testing.T) {
	fixtures := []fixture{
		{
			articles:     makeArticleFeed(5),
			marketing:    makeMarketingFeed(1),
			expectedFeed: []string{"Article", "Article", "Article", "Article", "Article", "ContentMarketing"},
		},
		{
			articles:  makeArticleFeed(10),
			marketing: makeMarketingFeed(2),
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
		{
			articles:  makeArticleFeed(10),
			marketing: makeMarketingFeed(1),
			expectedFeed: []string{
				"Article", "Article", "Article", "Article", "Article", "ContentMarketing",
				"Article", "Article", "Article", "Article", "Article", "Ads",
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

func makeArticleFeed(n int) []content.Article {
	feed := make([]content.Article, n)
	for i := range feed {
		a := content.NewArticle()
		a.Title = strconv.Itoa(i)
		feed[i] = a
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
