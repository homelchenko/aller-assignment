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
