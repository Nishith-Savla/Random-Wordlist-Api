package service

import (
	"github.com/Nishith-Savla/Random-Wordlist-Api/domain"
	"github.com/Nishith-Savla/Random-Wordlist-Api/dto"
)

type WordlistService interface {
	GetWords(limit int) *dto.WordlistResponse
	Shuffle()
}

type DefaultWordlistService struct {
	repo domain.WordlistRepository
}

func (s *DefaultWordlistService) GetWords(limit int) *dto.WordlistResponse {
	words := s.repo.GetWords(limit)
	return &dto.WordlistResponse{Words: words}
}

func (s *DefaultWordlistService) Shuffle() {
	s.repo.Shuffle()
}

func NewDefaultWordlistService(repo domain.WordlistRepository) *DefaultWordlistService {
	return &DefaultWordlistService{repo: repo}
}
