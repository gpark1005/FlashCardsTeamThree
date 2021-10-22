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

func (s Service) CreateMatching(f entities.Matching) (*repo.Database, error) {
	db, err := s.Repo.CreateMatching(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}

func (s Service) CreateInfoOnly(f entities.InfoOnly) (*repo.Database, error) {
	db, err := s.Repo.CreateFlashcard(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}

func (s Service) CreateQAndA(f entities.QAndA) (*repo.Database, error){
	db, err := s.Repo.CreateFlashcard(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}

func (s Service) CreateTOrF(f entities.TOrF) (*repo.Database, error){
	db, err := s.Repo.CreateFlashcard(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}

func (s Service) CreateMultipleChoice(f entities.MultipleChoice) (*repo.Database, error){
	db, err := s.Repo.CreateFlashcard(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}