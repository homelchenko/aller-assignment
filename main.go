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

	getNewsFeed(w, req)
}

func getNewsFeed(w http.ResponseWriter, req *http.Request) {
	ca := make(chan articleResult)
	cm := make(chan marketingResult)

	go downloadArticles(req.Context(), ca)
	go downloadMarketing(req.Context(), cm)

	ar := <-ca
	mr := <-cm

	if ar.handleError(w) || mr.handleError(w) {
		return
	}

	feed := news.ProduceNewsFeed(ar.articles, mr.marketing)

	encodeToJSON(w, feed)
}

func downloadArticles(ctx context.Context, c chan<- articleResult) {
	d := downloader.NewArticleFeedReader()

	as, err := d.Download(ctx)
	c <- articleResult{as, err}
}

type articleResult struct {
	articles []news.Article
	err      error
}

func (r articleResult) handleError(w http.ResponseWriter) bool {
	return handleDownloadingError(w, r.err)
}

func downloadMarketing(ctx context.Context, c chan<- marketingResult) {
	d := downloader.NewMarketingFeedReader()

	ms, err := d.Download(ctx)
	c <- marketingResult{ms, err}
}

type marketingResult struct {
	marketing []news.Marketing
	err       error
}

func (r marketingResult) handleError(w http.ResponseWriter) bool {
	return handleDownloadingError(w, r.err)
}

func handleDownloadingError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	log.Printf("error downloading feed: %s", err)
	w.WriteHeader(http.StatusInternalServerError)

	return true
}

func encodeToJSON(w http.ResponseWriter, feed news.Feed) {
	enc := json.NewEncoder(w)

	err := enc.Encode(feed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
