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

type Piece interface {
}

type Feed []Piece

func ProduceNewsFeed(a []Article, m []Marketing) Feed {
	pieces := []Piece{}

	for i, p := range a {
		pieces = append(pieces, p)

		if (i+1)%articleLen == 0 {
			mi := i / articleLen
			if mi < len(m) {
				pieces = append(pieces, m[mi])
			} else {
				ad := NewAd()
				pieces = append(pieces, ad)
			}
		}
	}

	return pieces
}
