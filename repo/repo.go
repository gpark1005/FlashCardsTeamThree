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

	for _, v := range fcSlice.Matching {
		if v
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

//func (r Repo) GetById(id string) (interface{}, error) {
//	var db []map[string]interface{}
//
//	file, err := ioutil.ReadFile(r.Filename)
//	if err != nil {
//		return nil, err
//	}
//
//	err = json.Unmarshal(file, &db)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, v := range db {
//		if idCheck, ok := db["Id"]; ok {
//			switch idCheck {
//			case idCheck == id:
//			}
//		}
//	}
//	return nil, err
//}










//	switch  {
//	case db.Matching:
//		if db.Matching == typ
//	}
//}
//		_, v := range db.Matching{
//		if v == id {
//			idFound = v
//			return nil, err
//		}
//	}
//	return nil, err
//
//}