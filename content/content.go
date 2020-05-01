package content

type Article struct {
	Type          string  `json:"string"`
	Harvester     string  `json:"harvesterId"`
	CerebroScore  float64 `json:"cerebro-score"`
	URL           string  `json:"url"`
	Title         string  `json:"title"`
	CleanImageURL string  `json:"cleanImage"`
}

type Marketing struct {
	Type          string  `json:"type"`
	Harvester     string  `json:"harvesterId"`
	Partner       string  `json:"commercialPartner"`
	LogoURL       string  `json:"logoURL"`
	CerebroScore  float64 `json:"cerebro-score"`
	URL           string  `json:"url"`
	Title         string  `json:"title"`
	CleanImageURL string  `json:"cleanImage"`
}

type Ad struct {
	Type string `json:"type"`
}

type NewsPiece interface {
}

type NewsFeed struct {
	Items []NewsPiece
}

func ProduceNewsFeed(a []Article, m []Marketing) NewsFeed {
	pieces := []NewsPiece{}
	for i, p := range a {
		pieces = append(pieces, p)
		if (i+1)%5 == 0 {
			pieces = append(pieces, m[0])
		}
	}

	return NewsFeed{Items: pieces}
}
