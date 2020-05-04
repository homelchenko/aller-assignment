package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/homelchenko/aller-assignment/news"
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

	feed := news.Feed{}
	enc := json.NewEncoder(w)

	err := enc.Encode(feed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
