package entities

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

func (m *Matching) UpdateById(id string, match Matching) error{
	file, err := ioutil.ReadFile("matching.json")
	if err != nil{
		return err
	}

	md := MatchingDatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.Matching{
		if v.Id == id{
			match.Id = id
			md.Matching[i] = match
		}
	}
	marshal, err := json.MarshalIndent(md.Matching, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile("matching.json", marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}