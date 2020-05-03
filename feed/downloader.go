package feed

import "context"

type MarketingFeedDownloader interface {
	Download(ctx context.Context) ([]Marketing, error)
}

type ArticleFeedDownloader interface {
	Download(ctx context.Context) ([]Article, error)
}
