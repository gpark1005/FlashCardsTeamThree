package service

import (
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"github.com/gpark1005/FlashCardsTeamThree/repo"
)

type Repo interface {
	CreateFlashcard (f entities.Flashcards) (*repo.Database, error)
}

type Service struct {
	Repo Repo
}

func NewService(r Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) CreateFlashcard (f entities.Flashcards) (*repo.Database, error) {
	db, err := s.Repo.CreateFlashcard(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}