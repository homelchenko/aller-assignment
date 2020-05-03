package news

const (
	articleLen = 5
)

type Article struct {
	Type          string  `json:"type"`
	Harvester     string  `json:"harvesterId"`
	CerebroScore  float64 `json:"cerebro-score"`
	URL           string  `json:"url"`
	Title         string  `json:"title"`
	CleanImageURL string  `json:"cleanImage"`
}

func NewArticle() Article {
	return Article{
		Type: "Article",
	}
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

func NewMarketing() Marketing {
	return Marketing{
		Type: "ContentMarketing",
	}
}

type Ad struct {
	Type string `json:"type"`
}

func NewAd() Ad {
	return Ad{Type: "Ads"}
}
