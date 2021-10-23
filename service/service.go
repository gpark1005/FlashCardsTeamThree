package service

import (
	"github.com/gpark1005/FlashCardsTeamThree/entities"
)

type Repo interface {
	CreateMatching(f entities.Matching) error
	CreateInfoOnly(f entities.InfoOnly) error
	CreateQAndA(f entities.QAndA) error
	CreateTOrF(f entities.TOrF) error
	CreateMultipleChoice(f entities.MultipleChoice) error
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
	err := s.Repo.CreateMatching(f)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CreateInfoOnly(f entities.InfoOnly) error {
	err := s.Repo.CreateInfoOnly(f)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CreateQAndA(f entities.QAndA) error {
	err := s.Repo.CreateQAndA(f)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CreateTOrF(f entities.TOrF) error {
	err := s.Repo.CreateTOrF(f)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CreateMultipleChoice(f entities.MultipleChoice) error {
	err := s.Repo.CreateMultipleChoice(f)
	if err != nil {
		return err
	}
	return nil
}
