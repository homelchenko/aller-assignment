package news

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
