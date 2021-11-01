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
	GetAllFlashcards() (*repo.Flashcards, error)
	GetMatchingById(id string) (*entities.Matching, error)
	GetInfoOnlyById(id string) (*entities.InfoOnly, error)
	GetQAndAById(id string) (*entities.QAndA, error)
	GetTOrFById(id string) (*entities.TOrF, error)
	GetMultipleChoiceById(id string) (*entities.MultipleChoice, error)
	UpdateMatchingById(id string, m entities.Matching) error
	UpdateInfoById(id string, io entities.InfoOnly) error
	UpdateQandAById(id string, qa entities.QAndA) error
	UpdateMultipleChoiceById(id string, mc entities.MultipleChoice) error
	UpdateTorFById(id string, tf entities.TOrF) error
	DeleteMatchingById (id string) error
	DeleteInfoOnlyById (id string) error
	DeleteTOrFById (id string) error
	DeleteQAndAById (id string) error
	DeleteMCById (id string) error
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


func (s Service) GetAllFlashcards() (*repo.Flashcards, error) {
	fcSlice, err := s.Repo.GetAllFlashcards()
	if err != nil {
		return nil, err
	}
	return fcSlice, nil
}

func (s Service) GetMatchingById(id string) (*entities.Matching, error) {
	fc, err := s.Repo.GetMatchingById(id)
	if err != nil{
		return nil, err
	}
	return fc, nil
}

func (s Service) GetInfoOnlyById(id string) (*entities.InfoOnly, error){
	fc, err := s.Repo.GetInfoOnlyById(id)
	if err != nil{
		return nil, err
	}
	return fc, nil
}

func (s Service) GetQAndAById(id string) (*entities.QAndA, error){
	fc, err := s.Repo.GetQAndAById(id)
	if err != nil{
		return nil, err
	}
	return fc, nil
}

func (s Service) GetTOrFById(id string) (*entities.TOrF, error){
	fc, err := s.Repo.GetTOrFById(id)
	if err != nil{
		return nil, err
	}
	return fc, nil
}

func (s Service) GetMultipleChoiceById(id string) (*entities.MultipleChoice, error){
	fc, err := s.Repo.GetMultipleChoiceById(id)
	if err != nil{
		return nil, err
	}
	return fc, nil
}

func (s Service) UpdateByID(id string, m entities.Matching) error{
	err := s.Repo.UpdateMatchingById(id, m)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) UpdateMatchingById(id string, m entities.Matching) error{
	err := s.Repo.UpdateMatchingById(id, m)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) UpdateInfoById(id string, io entities.InfoOnly) error{
	err := s.Repo.UpdateInfoById(id, io)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) UpdateQandAById(id string, qa entities.QAndA) error {
	err := s.Repo.UpdateQandAById(id, qa)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) UpdateMultipleChoiceById(id string, mc entities.MultipleChoice) error{
	err := s.Repo.UpdateMultipleChoiceById(id, mc)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) UpdateTorFById(id string, tf entities.TOrF) error {
	err := s.Repo.UpdateTorFById(id, tf)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) DeleteMatchingById(id string)error{
	err := s.Repo.DeleteMatchingById(id)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) DeleteInfoOnlyById(id string)error{
	err := s.Repo.DeleteInfoOnlyById(id)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) DeleteQAndAById(id string)error{
	err := s.Repo.DeleteQAndAById(id)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) DeleteTOrFById(id string)error{
	err := s.Repo.DeleteTOrFById(id)
	if err != nil{
		return err
	}
	return nil
}

func (s Service) DeleteMCById(id string)error{
	err := s.Repo.DeleteMCById(id)
	if err != nil{
		return err
	}
	return nil
}