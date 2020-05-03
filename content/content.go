package content

type Article struct {
	Type          string  `json:"type"`
	Harvester     string  `json:"harvesterId"`
	CerebroScore  float64 `json:"cerebro-score"`
	URL           string  `json:"url"`
	Title         string  `json:"title"`
	CleanImageURL string  `json:"cleanImage"`
}

func (a *Article) PieceType() string {
	return a.Type
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

func (m *Marketing) PieceType() string {
	return m.Type
}

type Ad struct {
	Type string `json:"type"`
}

func (a *Ad) PieceType() string {
	return a.Type
}

type NewsPiece interface {
	PieceType() string
}

type NewsFeed struct {
	Items []NewsPiece
}

func ProduceNewsFeed(a []Article, m []Marketing) NewsFeed {
	pieces := []NewsPiece{}

	for i, p := range a {
		pc := p
		pieces = append(pieces, &pc)

		if (i+1)%5 == 0 {
			mi := 18
			if mi < len(m) {
				pieces = append(pieces, &m[0])
			} else {
				pieces = append(pieces, &Ad{Type: "Ads"})
			}
		}
	}

	return NewsFeed{Items: pieces}
}
