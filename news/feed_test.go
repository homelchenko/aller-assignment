package news_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/homelchenko/aller-assignment/news"
)

type fixture struct {
	articles     []news.Article
	marketing    []news.Marketing
	expectedFeed []news.Piece
}

func TestProduceNewsFeedForEmptySlices(t *testing.T) {
	fixtures := []fixture{
		{
			articles:     nil,
			marketing:    nil,
			expectedFeed: []news.Piece{},
		},
		{
			articles:     makeArticleFeed(0),
			marketing:    nil,
			expectedFeed: []news.Piece{},
		},
		{
			articles:     nil,
			marketing:    makeMarketingFeed(0),
			expectedFeed: []news.Piece{},
		},
		{
			articles:     makeArticleFeed(0),
			marketing:    makeMarketingFeed(0),
			expectedFeed: []news.Piece{},
		},
	}

	for _, fixture := range fixtures {
		feed := news.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(feed), len(fixture.expectedFeed))
		}
	}
}

func TestProduceNewsFeedWhenArticlesFewerThanFive(t *testing.T) {
	articles := makeArticleFeed(4)
	fixtures := []fixture{
		{
			articles:     articles,
			marketing:    nil,
			expectedFeed: []news.Piece{articles[0], articles[1], articles[2], articles[3]},
		},
		{
			articles:     articles,
			marketing:    makeMarketingFeed(1),
			expectedFeed: []news.Piece{articles[0], articles[1], articles[2], articles[3]},
		},
	}

	for _, fixture := range fixtures {
		feed := news.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(feed), len(fixture.expectedFeed))
		}

		for i, item := range feed {
			if !reflect.DeepEqual(item, fixture.expectedFeed[i]) {
				t.Errorf("At %d got %s, but expected %s", i, item, fixture.expectedFeed[i])
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
		expectedFeed: []news.Piece{
			articles[0], articles[1], articles[2], articles[3], articles[4], marketing[0],
		},
	}
	fixtures = append(fixtures, fix)

	articles = makeArticleFeed(10)
	marketing = makeMarketingFeed(2)
	fix = fixture{
		articles:  articles,
		marketing: marketing,
		expectedFeed: []news.Piece{
			articles[0], articles[1], articles[2], articles[3], articles[4], marketing[0],
			articles[5], articles[6], articles[7], articles[8], articles[9], marketing[1],
		},
	}
	fixtures = append(fixtures, fix)

	for _, fixture := range fixtures {
		feed := news.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(feed), len(fixture.expectedFeed))
		}

		for i, item := range feed {
			if !reflect.DeepEqual(item, fixture.expectedFeed[i]) {
				t.Errorf("At %d got %s, but expected %s", i, item, fixture.expectedFeed[i])
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
		expectedFeed: []news.Piece{
			articles[0], articles[1], articles[2], articles[3], articles[4], ads[0],
		},
	}
	fixtures = append(fixtures, fix)

	articles = makeArticleFeed(10)
	marketing := makeMarketingFeed(1)
	fix = fixture{
		articles:  articles,
		marketing: marketing,
		expectedFeed: []news.Piece{
			articles[0], articles[1], articles[2], articles[3], articles[4], marketing[0],
			articles[5], articles[6], articles[7], articles[8], articles[9], ads[0],
		},
	}
	fixtures = append(fixtures, fix)

	for _, fixture := range fixtures {
		feed := news.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(feed) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(feed), len(fixture.expectedFeed))
		}

		for i, item := range feed {
			if !reflect.DeepEqual(item, fixture.expectedFeed[i]) {
				t.Errorf("At %d got %s, but expected %s", i, item, fixture.expectedFeed[i])
				break
			}
		}
	}
}

func makeArticleFeed(n int) []news.Article {
	feed := make([]news.Article, n)
	for i := range feed {
		a := news.NewArticle()
		a.Title = strconv.Itoa(i)
		feed[i] = a
	}

	return feed
}

func makeMarketingFeed(n int) []news.Marketing {
	feed := make([]news.Marketing, n)
	for i := range feed {
		m := news.NewMarketing()
		m.Title = strconv.Itoa(i)
		feed[i] = m
	}

	return feed
}

func makeAdFeed(n int) []news.Ad {
	feed := make([]news.Ad, n)
	for i := range feed {
		feed[i] = news.NewAd()
	}

	return feed
}
