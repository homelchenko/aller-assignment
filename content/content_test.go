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

		if len(feed.Items) != len(fixture.expectedNews) {
			t.Errorf("Got %d, but expected %d", len(feed.Items), len(fixture.expectedNews))
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
	var fixtures []fixture

	articles := makeArticleFeed(5)
	marketing := makeMarketingFeed(1)
	fix := fixture{
		articles:  articles,
		marketing: marketing,
		expectedNews: []content.NewsPiece{
			&articles[0], &articles[1], &articles[2], &articles[3], &articles[4], &marketing[0],
		},
	}
	fixtures = append(fixtures, fix)

	articles = makeArticleFeed(10)
	marketing = makeMarketingFeed(2)
	fix = fixture{
		articles:  articles,
		marketing: marketing,
		expectedNews: []content.NewsPiece{
			&articles[0], &articles[1], &articles[2], &articles[3], &articles[4], &marketing[0],
			&articles[5], &articles[6], &articles[7], &articles[8], &articles[9], &marketing[1],
		},
	}
	fixtures = append(fixtures, fix)

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

func TestProduceNewsFeedWhenNotEnoughMarketingForEveryFiveArticles(t *testing.T) {
	var fixtures []fixture

	articles := makeArticleFeed(5)
	ads := makeAdFeed(1)
	fix := fixture{
		articles:  articles,
		marketing: makeMarketingFeed(0),
		expectedNews: []content.NewsPiece{
			&articles[0], &articles[1], &articles[2], &articles[3], &articles[4], &ads[0],
		},
	}
	fixtures = append(fixtures, fix)

	articles = makeArticleFeed(10)
	marketing := makeMarketingFeed(1)
	fix = fixture{
		articles:  articles,
		marketing: marketing,
		expectedNews: []content.NewsPiece{
			&articles[0], &articles[1], &articles[2], &articles[3], &articles[4], &marketing[0],
			&articles[5], &articles[6], &articles[7], &articles[8], &articles[9], &ads[0],
		},
	}
	fixtures = append(fixtures, fix)

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
		m := content.NewMarketing()
		m.Title = strconv.Itoa(i)
		feed[i] = m
	}

	return feed
}

func makeAdFeed(n int) []content.Ad {
	feed := make([]content.Ad, n)
	for i := range feed {
		feed[i] = content.NewAd()
	}

	return feed
}
