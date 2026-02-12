package shorten

type ShortenService struct {
}

func NewShortenService() *ShortenService {
	return &ShortenService{}
}

func (s *ShortenService) ShortenURL(longURL string) (string, error) {
}
