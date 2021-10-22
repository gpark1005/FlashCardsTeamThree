package repo

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"io/ioutil"
)

type Database struct {
	Database []entities.Flashcards
}

type Repo struct {
	Filename string
}

func NewRepo(f string) Repo{
	return Repo{
		Filename: f,
	}
}

func (r Repo) CreateMatching(f entities.Matching) (*Database, error) {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil{
		return nil, err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil{
		return nil, err
	}

	fcSlice.Database = append(fcSlice.Database, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil{
		return nil, err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil{
		return nil, err
	}

	return &fcSlice, nil
}

func (r Repo) CreateInfoOnly(f entities.InfoOnly) (*Database, error) {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil{
		return nil, err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil{
		return nil, err
	}

	fcSlice.Database = append(fcSlice.Database, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil{
		return nil, err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil{
		return nil, err
	}

	return &fcSlice, nil
}

func (r Repo) CreateQAndA(f entities.QAndA) (*Database, error) {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil{
		return nil, err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil{
		return nil, err
	}

	fcSlice.Database = append(fcSlice.Database, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil{
		return nil, err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil{
		return nil, err
	}

	return &fcSlice, nil
}

func (r Repo) CreateTOrF(f entities.TOrF) (*Database, error) {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil{
		return nil, err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil{
		return nil, err
	}

	fcSlice.Database = append(fcSlice.Database, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil{
		return nil, err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil{
		return nil, err
	}

	return &fcSlice, nil
}

func (r Repo) CreateMultipleChoice(f entities.MultipleChoice) (*Database, error) {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil{
		return nil, err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil{
		return nil, err
	}

	fcSlice.Database = append(fcSlice.Database, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil{
		return nil, err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil{
		return nil, err
	}

	return &fcSlice, nil
}