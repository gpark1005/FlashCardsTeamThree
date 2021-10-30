package entities

import (
	"errors"
)

type MatchingDatabase struct {
	Matching       []Matching
}

type Matching struct {
	Id       string
	Ctype    string                 `json:"type"`
	Category string                 `json:"category"`
	Question map[string]interface{} `json:"question"`
	Choices  map[string]interface{} `json:"choices"`
	Answers  map[string]interface{} `json:"answers"`
}

func (m Matching) ValidateMatching() error {
	if m.Category == ""{
		return errors.New("category required")
	} else if m.Answers == nil{
		return errors.New("answer key required")
	}else if m.Choices == nil{
		return errors.New("choices required")
	}else if m.Question == nil{
		return errors.New("questions required")
	}
	return nil
}

