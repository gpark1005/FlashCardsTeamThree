package service

import (
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"github.com/google/uuid"
)

type Repo interface {
	CreateFlashcard (f entities.Flashcards) error
}

type Service struct {
	Repo Repo
}

func NewService(r Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) CreateFlashcard (f entities.Flashcards) error {
	s.Repo.CreateFlashcard(f)
}