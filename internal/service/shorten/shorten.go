package shorten

import (
	"context"
	"fmt"
	"urlshorter/internal/config"
	"urlshorter/internal/repositories"
	gendigit "urlshorter/pkgs/gen_digit"
)

type ShortenService struct {
	cfg        *config.Config
	repository *repositories.ShortenRepository
}

func NewShortenService(config *config.Config) *ShortenService {
	return &ShortenService{
		cfg:        config,
		repository: repositories.NewShortenRepository(),
	}
}

func (s *ShortenService) ShortenURL(ctx context.Context, longURL string) (string, error) {
	gd, err := gendigit.GenerateShortCode(s.cfg.Charset, 6)
	if err != nil {
		return "", err
	}
	if gd == "" {
		return "", nil
	}

	if err := s.repository.Save(ctx, gd, longURL); err != nil {
		return "", err
	}

	return fmt.Sprintf("http://localhost:%s/%s", s.cfg.Port, gd), nil
}
