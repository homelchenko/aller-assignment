package news_test

import (
	"testing"

	"github.com/homelchenko/aller-assignment/news"
)

func TestNewArticle(t *testing.T) {
	a := news.NewArticle()

	if a.Type != "Article" {
		t.Errorf("Got '%s', but expected 'Article'", a.Type)
	}
}

func TestNewMarketing(t *testing.T) {
	m := news.NewMarketing()

	if m.Type != "ContentMarketing" {
		t.Errorf("Got %s, but expected 'ContentMarketing'", m.Type)
	}
}

func TestNewAd(t *testing.T) {
	a := news.NewAd()

	if a.Type != "Ads" {
		t.Errorf("Got %s, but expected 'Ads'", a.Type)
	}
}
