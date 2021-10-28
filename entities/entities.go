package entities

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type FcType struct {
	Type string `json:"Type"`
}

type Flashcards struct {
	matching       []Matching
	infoOnly       InfoOnly
	qAndA          QAndA
	tOrF           TOrF
	multipleChoice MultipleChoice
}

type Matching struct {
	Id       string
	Ctype    string                 `json:"type"`
	Category string                 `json:"category"`
	Question map[string]interface{} `json:"question"`
	Choices  map[string]interface{} `json:"choices"`
	Answers  map[string]interface{} `json:"answers"`
}

type InfoOnly struct {
	Id          string
	Ctype       string `json:"type"`
	Category 	string `json:"category"`
	Information string `json:"information"`
}

type QAndA struct {
	Id       string
	Ctype    string `json:"type"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type TOrF struct {
	Id       string
	Ctype    string `json:"type"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type MultipleChoice struct {
	Id       string
	Ctype    string                 `json:"type"`
	Category string                 `json:"category"`
	Question string                 `json:"question"`
	Choices  map[string]interface{} `json:"choices"`
	Answer   string                 `json:"answer"`
}

func ValidateType(t FcType) (bool, error) {
	fileType := struct{
		Types []string
	}{}

	data, err := ioutil.ReadFile("types.json")
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &fileType)

	for _, v := range fileType.Types {
		if t.Type == v{
			return true, nil
		}

	}
	return false, errors.New("type not valid")
}

func (m Matching) ValidateMatching() error {
	if m.Category == ""{
		return errors.New("category required")
	} else if m.Answers == nil{
		return errors.New("answer key required")
	}else if m.Choices == nil{
		return errors.New("choices required")
	}else if m.Question == nil{
		return errors.New("question required")
	}
	return nil
}


func (i InfoOnly) ValidateInfoOnly() error {
	if i.Category == ""{
		return errors.New("category required")
	}else if i.Information == "" {
		return errors.New("information required")
	}
	return nil
}

func (q QAndA) ValidateQAndA() error {
	if q.Category == ""{
		return errors.New("category required")
	}else if q.Question == "" {
		return errors.New("questions required")
	}else if q.Answer == "" {
		return errors.New("answer required")
	}
	return nil
}

func (t TOrF) ValidateTOrF() error {
	if t.Category == ""{
		return errors.New("category required")
	}else if t.Question == "" {
		return errors.New("questions required")
	}else if t.Answer == "" {
		return errors.New("answer required")
	}
	return nil
}

func (mc MultipleChoice) ValidateMultipleChoice() error {
	if mc.Category == ""{
		return errors.New("category required")
	}else if mc.Question == "" {
		return errors.New("questions required")
	}else if mc.Answer == "" {
		return errors.New("answer required")
	}else if mc.Choices == nil {
		return errors.New("choices required")
	}
	return nil
}