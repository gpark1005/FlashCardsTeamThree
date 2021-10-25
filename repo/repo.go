package repo

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"io/ioutil"
)

type Database struct {
	Flashcards []interface{}
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

	fcSlice.Flashcards = append(fcSlice.Flashcards, f)

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

	fcSlice.Flashcards = append(fcSlice.Flashcards, f)

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

	fcSlice.Flashcards = append(fcSlice.Flashcards, f)

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

	fcSlice.Flashcards = append(fcSlice.Flashcards, f)

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

	fcSlice.Flashcards = append(fcSlice.Flashcards, f)

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
//	var db, foundTypes Database
//	file, err := ioutil.ReadFile(r.Filename)
//	if err != nil{
//		return nil, err
//	}
//
//	err = json.Unmarshal(file, &db)
//	if err != nil{
//		return nil, err
//	}
//}