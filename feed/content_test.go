package feed_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/homelchenko/aller-assignment/feed"
)

type fixture struct {
	articles     []feed.Article
	marketing    []feed.Marketing
	expectedFeed []feed.NewsPiece
}

func TestProduceNewsFeedForEmptySlices(t *testing.T) {
	fixtures := []fixture{
		{
			articles:     nil,
			marketing:    nil,
			expectedFeed: []feed.NewsPiece{},
		},
		{
			articles:     makeArticleFeed(0),
			marketing:    nil,
			expectedFeed: []feed.NewsPiece{},
		},
		{
			articles:     nil,
			marketing:    makeMarketingFeed(0),
			expectedFeed: []feed.NewsPiece{},
		},
		{
			articles:     makeArticleFeed(0),
			marketing:    makeMarketingFeed(0),
			expectedFeed: []feed.NewsPiece{},
		},
	}

	for _, fixture := range fixtures {
		f := feed.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(f) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(f), len(fixture.expectedFeed))
		}
	}
}

func TestProduceNewsFeedWhenArticlesFewerThanFive(t *testing.T) {
	articles := makeArticleFeed(4)
	fixtures := []fixture{
		{
			articles:     articles,
			marketing:    nil,
			expectedFeed: []feed.NewsPiece{&articles[0], &articles[1], &articles[2], &articles[3]},
		},
		{
			articles:     articles,
			marketing:    makeMarketingFeed(1),
			expectedFeed: []feed.NewsPiece{&articles[0], &articles[1], &articles[2], &articles[3]},
		},
	}

	for _, fixture := range fixtures {
		f := feed.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(f) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(f), len(fixture.expectedFeed))
		}

		for i, item := range f {
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
		expectedFeed: []feed.NewsPiece{
			&articles[0], &articles[1], &articles[2], &articles[3], &articles[4], &marketing[0],
		},
	}
	fixtures = append(fixtures, fix)

	articles = makeArticleFeed(10)
	marketing = makeMarketingFeed(2)
	fix = fixture{
		articles:  articles,
		marketing: marketing,
		expectedFeed: []feed.NewsPiece{
			&articles[0], &articles[1], &articles[2], &articles[3], &articles[4], &marketing[0],
			&articles[5], &articles[6], &articles[7], &articles[8], &articles[9], &marketing[1],
		},
	}
	fixtures = append(fixtures, fix)

	for _, fixture := range fixtures {
		f := feed.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(f) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(f), len(fixture.expectedFeed))
		}

		for i, item := range f {
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
		expectedFeed: []feed.NewsPiece{
			&articles[0], &articles[1], &articles[2], &articles[3], &articles[4], &ads[0],
		},
	}
	fixtures = append(fixtures, fix)

	articles = makeArticleFeed(10)
	marketing := makeMarketingFeed(1)
	fix = fixture{
		articles:  articles,
		marketing: marketing,
		expectedFeed: []feed.NewsPiece{
			&articles[0], &articles[1], &articles[2], &articles[3], &articles[4], &marketing[0],
			&articles[5], &articles[6], &articles[7], &articles[8], &articles[9], &ads[0],
		},
	}
	fixtures = append(fixtures, fix)

	for _, fixture := range fixtures {
		f := feed.ProduceNewsFeed(fixture.articles, fixture.marketing)

		if len(f) != len(fixture.expectedFeed) {
			t.Errorf("Got %d, but expected %d", len(f), len(fixture.expectedFeed))
		}

		for i, item := range f {
			if !reflect.DeepEqual(item, fixture.expectedFeed[i]) {
				t.Errorf("At %d got %s, but expected %s", i, item, fixture.expectedFeed[i])
				break
			}
		}
	}
}

func makeArticleFeed(n int) []feed.Article {
	f := make([]feed.Article, n)
	for i := range f {
		a := feed.NewArticle()
		a.Title = strconv.Itoa(i)
		f[i] = a
	}

	return f
}

func makeMarketingFeed(n int) []feed.Marketing {
	f := make([]feed.Marketing, n)
	for i := range f {
		m := feed.NewMarketing()
		m.Title = strconv.Itoa(i)
		f[i] = m
	}

	return f
}

func makeAdFeed(n int) []feed.Ad {
	f := make([]feed.Ad, n)
	for i := range f {
		f[i] = feed.NewAd()
	}

	return f
}
