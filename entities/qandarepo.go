package entities

import "errors"

type QAndADatabase struct {
	QAndA      []QAndA
}

type QAndA struct {
	Id       string
	Ctype    string `json:"type"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (q QAndA) ValidateQAndA() error {
	if q.Category == ""{
		return errors.New("category required")
	}else if q.Question == "" {
		return errors.New("question(s) required")
	}else if q.Answer == "" {
		return errors.New("answer required")
	}
	return nil
}
