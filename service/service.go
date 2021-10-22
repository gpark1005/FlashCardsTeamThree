package service

import (
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"github.com/gpark1005/FlashCardsTeamThree/repo"
)

type Repo interface {
	CreateMatching (f entities.Matching) (*repo.Database, error)
	CreateInfoOnly(f entities.InfoOnly) (*repo.Database, error)
	CreateQAndA(f entities.QAndA) (*repo.Database, error)
	CreateTOrF(f entities.TOrF) (*repo.Database, error)
	CreateMultipleChoice(f entities.MultipleChoice) (*repo.Database, error)
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
	db, err := s.Repo.CreateInfoOnly(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}

func (s Service) CreateQAndA(f entities.QAndA) (*repo.Database, error){
	db, err := s.Repo.CreateQAndA(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}

func (s Service) CreateTOrF(f entities.TOrF) (*repo.Database, error){
	db, err := s.Repo.CreateTOrF(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}

func (s Service) CreateMultipleChoice(f entities.MultipleChoice) (*repo.Database, error){
	db, err := s.Repo.CreateMultipleChoice(f)
	if err != nil{
		return nil, err
	}
	return db, nil
}