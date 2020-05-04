package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/homelchenko/aller-assignment/news"
	"github.com/homelchenko/aller-assignment/news/downloader"
)

func main() {
	http.HandleFunc("/news", newsFeed)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func newsFeed(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	articleDowloader := downloader.NewArticleFeedReader()
	marketingDownloader := downloader.NewMarketingFeedReader()

	ca := make(chan articleResult)
	cm := make(chan marketingResult)

	go downloadArticles(req.Context(), ca, articleDowloader)
	go downloadMarketing(req.Context(), cm, marketingDownloader)

	ar := <-ca
	mr := <-cm

	if ar.err != nil {
		log.Printf("error downloading feed: %s", ar.err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if mr.err != nil {
		log.Printf("error downloading feed: %s", mr.err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	feed := news.ProduceNewsFeed(ar.articles, mr.marketing)
	enc := json.NewEncoder(w)

	err := enc.Encode(feed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func downloadArticles(ctx context.Context, c chan<- articleResult, d news.ArticleFeedDownloader) {
	as, err := d.Download(ctx)
	c <- articleResult{as, err}
}

type articleResult struct {
	articles []news.Article
	err      error
}

func downloadMarketing(ctx context.Context, c chan<- marketingResult, d news.MarketingFeedDownloader) {
	ms, err := d.Download(ctx)
	c <- marketingResult{ms, err}
}

type marketingResult struct {
	marketing []news.Marketing
	err       error
}
