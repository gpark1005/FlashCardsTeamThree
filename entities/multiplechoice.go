package entities

import "errors"

type MCDatabase struct {
	MultipleChoice       []MultipleChoice
}

type MultipleChoice struct {
	Id       string
	Ctype    string                 `json:"type"`
	Category string                 `json:"category"`
	Question string                 `json:"question"`
	Choices  map[string]interface{} `json:"choices"`
	Answer   string                 `json:"answer"`
}

func (mc MultipleChoice) ValidateMultipleChoice() error {
	if mc.Category == ""{
		return errors.New("category required")
	}else if mc.Question == "" {
		return errors.New("question(s) required")
	}else if mc.Answer == "" {
		return errors.New("answer required")
	}else if mc.Choices == nil {
		return errors.New("choices required")
	}
	return nil
}