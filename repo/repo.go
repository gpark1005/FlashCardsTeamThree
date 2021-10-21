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

func (r Repo) CreateFlashcard(f entities.Flashcards) (*Database, error) {
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
