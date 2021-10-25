package repo

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"io/ioutil"
)

type Database struct {
	Matching       []entities.Matching
	InfoOnly       []entities.InfoOnly
	QAndA          []entities.QAndA
	TOrF           []entities.TOrF
	MultipleChoice []entities.MultipleChoice
}


type Data struct {
	Data interface{}
}

type Repo struct {
	Filename string
}

func NewRepo(f string) Repo {
	return Repo{
		Filename: f,
	}
}

func (r Repo) CreateMatching(f entities.Matching) error {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.Matching = append(fcSlice.Matching, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateInfoOnly(f entities.InfoOnly) error {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.InfoOnly = append(fcSlice.InfoOnly, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateQAndA(f entities.QAndA) error {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.QAndA = append(fcSlice.QAndA, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateTOrF(f entities.TOrF) error {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.TOrF = append(fcSlice.TOrF, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateMultipleChoice(f entities.MultipleChoice) error {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.MultipleChoice = append(fcSlice.MultipleChoice, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}


func (r Repo) GetAllFlashcards() (*Database, error) {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return nil, err
	}
	return &fcSlice, nil
}

//func (r Repo) GetByType(ct string) (*Database, error){
//	db := Database{}
//	var foundTypes []interface{}
//
//	file, err := ioutil.ReadFile(r.Filename)
//	if err != nil{
//		return nil, err
//	}
//
//	err = json.Unmarshal(file, &db)
//	if err != nil{
//		return nil, err
//	}
//	for _, v := range db.Flashcards{
//		foundTypes = append(foundTypes, v)
//	}
//
//}