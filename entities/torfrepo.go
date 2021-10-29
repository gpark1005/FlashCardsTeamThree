package entities

import "errors"

type TOrFDatabase struct {
	TOrF      []TOrF
}

type TOrF struct {
	Id       string
	Ctype    string `json:"type"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (t TOrF) ValidateTOrF() error {
	if t.Category == ""{
		return errors.New("category required")
	}else if t.Question == "" {
		return errors.New("question(s) required")
	}else if t.Answer == "" {
		return errors.New("answer required")
	}
	return nil
}