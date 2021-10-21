package repo

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"io/ioutil"
)

type Database struct {
	Database []entities.FlashCards
}

type Repo struct {
	Filename string
}

func NewRepo(f string) Repo{
	return Repo{
		Filename: f,
	}
}

func (r Repo) CreateFlashCard(f entities.FlashCards) error {
	fcSlice := Database{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil{
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil{
		return err
	}

	fcSlice.Database = append(fcSlice.Database, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil{
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil{
		return err
	}

	return nil
}
