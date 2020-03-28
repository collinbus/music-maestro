package domain

type Urls struct {
	Internal string
	External string
	Uri      string
}

func NewUrls(internal string, external string, uri string) *Urls {
	return &Urls{Internal: internal, External: external, Uri: uri}
}
