package repositories

import "context"

type ShortenRepository struct {
}

func NewShortenRepository() *ShortenRepository {
	return &ShortenRepository{}
}

func (r *ShortenRepository) Save(ctx context.Context, code string, longURL string) error {
	return nil
}
