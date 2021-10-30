package repo

import (
	"encoding/json"
	"errors"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"reflect"

	//"github.com/gpark1005/FlashCardsTeamThree/infoOnlyRepo"
	//"github.com/gpark1005/FlashCardsTeamThree/matchingRepo"
	//"github.com/gpark1005/FlashCardsTeamThree/multiplechoiceRepo"
	//"github.com/gpark1005/FlashCardsTeamThree/qandaRepo"
	//"github.com/gpark1005/FlashCardsTeamThree/torfRepo"
	"io/ioutil"
)

type Database struct {
	Flashcards []interface{}
}

type flashcards struct {
	Flashcards []map[string]interface{}
}

type Repo struct {
	MatchingFile string
	InfoOnlyFile string
	QAndAFile string
	MultipleChoiceFile string
	TrueOrFalseFile string
}

func NewRepo(m string, io string, mc string, tf string, qa string) Repo {
	return Repo{
		MatchingFile: m,
		InfoOnlyFile: io,
		QAndAFile: qa,
		MultipleChoiceFile: mc,
		TrueOrFalseFile: tf,
	}
}

func (r Repo) CreateMatching(f entities.Matching) error {
	fcSlice := entities.MatchingDatabase{}

	file, err := ioutil.ReadFile(r.MatchingFile)
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

	err = ioutil.WriteFile(r.MatchingFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateInfoOnly(f entities.InfoOnly) error {
	fcSlice := entities.InfoOnlyDatabase{}

	file, err := ioutil.ReadFile(r.InfoOnlyFile)
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

	err = ioutil.WriteFile(r.InfoOnlyFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateQAndA(f entities.QAndA) error {
	fcSlice := entities.QAndADatabase{}

	file, err := ioutil.ReadFile(r.QAndAFile)
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

	err = ioutil.WriteFile(r.QAndAFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateTOrF(f entities.TOrF) error {
	fcSlice := entities.TOrFDatabase{}

	file, err := ioutil.ReadFile(r.TrueOrFalseFile)
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

	err = ioutil.WriteFile(r.TrueOrFalseFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateMultipleChoice(f entities.MultipleChoice) error {
	fcSlice := entities.MCDatabase{}

	file, err := ioutil.ReadFile(r.MultipleChoiceFile)
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

	err = ioutil.WriteFile(r.MultipleChoiceFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}


func (r Repo) GetAllFlashcards() (*Database, error) {
	fcSlice := Database{}
	files := []string{r.MatchingFile, r.MultipleChoiceFile, r.TrueOrFalseFile, r.QAndAFile, r.InfoOnlyFile}

	for _, v := range files {
		file, err := ioutil.ReadFile(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(file, &fcSlice)
		if err != nil {
			return nil, err
		}
	}
	return &fcSlice, nil
}

func (r Repo) GetById(id string) (map[string]interface{}, error) {
	db := flashcards{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &db)
	if err != nil {
		return nil, err
	}

	for _, v := range db.Flashcards{
		if idCheck, ok := v["Id"]; ok{
			if idCheck == id{
				return v, nil
			}
		}
	}
	return nil, errors.New("flashcard does not exist")
}

func (r Repo) UpdateMatchingById(id string, m entities.Matching) error {
	file, err := ioutil.ReadFile(r.MatchingFile)
	if err != nil{
		return err
	}

	md := entities.MatchingDatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.Matching{
		if v.Id == id{
			m.Id = id
			md.Matching[i] = m
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.MatchingFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) UpdateInfoById(id string, io entities.InfoOnly) error{
	file, err := ioutil.ReadFile(r.InfoOnlyFile)
	if err != nil{
		return err
	}

	md := entities.InfoOnlyDatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.InfoOnly{
		if v.Id == id{
			io.Id = id
			md.InfoOnly[i] = io
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.InfoOnlyFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) UpdateQandAById(id string, qa entities.QAndA) error{
	file, err := ioutil.ReadFile(r.QAndAFile)
	if err != nil{
		return err
	}

	md := entities.QAndADatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.QAndA{
		if v.Id == id{
			qa.Id = id
			md.QAndA[i] = qa
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.QAndAFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) UpdateMultipleChoiceById(id string, mc entities.MultipleChoice) error{
	file, err := ioutil.ReadFile(r.MultipleChoiceFile)
	if err != nil{
		return err
	}

	md := entities.MCDatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.MultipleChoice{
		if v.Id == id{
			mc.Id = id
			md.MultipleChoice[i] = mc
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.MultipleChoiceFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) UpdateTorFById(id string, tf entities.TOrF) error {
	file, err := ioutil.ReadFile(r.TrueOrFalseFile)
	if err != nil{
		return err
	}

	md := entities.TOrFDatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.TOrF{
		if v.Id == id{
			tf.Id = id
			md.TOrF[i] = tf
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.TrueOrFalseFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}
