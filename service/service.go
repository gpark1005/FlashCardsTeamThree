package service

import (
	"github.com/google/uuid"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"github.com/gpark1005/FlashCardsTeamThree/repo"
)

type Repo interface {
	CreateMatching(f entities.Matching) error
	CreateInfoOnly(f entities.InfoOnly) error
	CreateQAndA(f entities.QAndA) error
	CreateTOrF(f entities.TOrF) error
	CreateMultipleChoice(f entities.MultipleChoice) error
	GetAllFlashcards() (*repo.Database, error)
	//GetById(id string) (interface{}, error)
}

type Service struct {
	Repo Repo
}

func NewService(r Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) CreateMatching(f entities.Matching) error {
	f.Id = uuid.New().String()

	err := s.Repo.CreateMatching(f)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CreateInfoOnly(f entities.InfoOnly) error {
	f.Id = uuid.New().String()

	err := s.Repo.CreateInfoOnly(f)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CreateQAndA(f entities.QAndA) error {
	f.Id = uuid.New().String()

	err := s.Repo.CreateQAndA(f)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CreateTOrF(f entities.TOrF) error {
	f.Id = uuid.New().String()

	err := s.Repo.CreateTOrF(f)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CreateMultipleChoice(f entities.MultipleChoice) error {
	f.Id = uuid.New().String()

	err := s.Repo.CreateMultipleChoice(f)
	if err != nil {
		return err
	}
	return nil
}


func (s Service) GetAllFlashcards() (*repo.Database, error) {
	fcSlice, err := s.Repo.GetAllFlashcards()
	if err != nil {
		return nil, err
	}
	return fcSlice, nil
}

//func (s Service) GetById(id string) (interface{}, error) {
//	fc, err := s.Repo.GetById(id)
//	if err != nil{
//		return nil, err
//	}
//	return fc, nil
//}