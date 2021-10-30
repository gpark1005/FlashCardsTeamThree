package entities

import "errors"

type InfoOnlyDatabase struct {
	InfoOnly       []InfoOnly
}

type InfoOnly struct {
	Id          string
	Ctype       string `json:"type"`
	Category 	string `json:"category"`
	Information string `json:"information"`
}

func (i InfoOnly) ValidateInfoOnly() error {
	if i.Category == ""{
		return errors.New("category required")
	}else if i.Information == "" {
		return errors.New("information required")
	}
	return nil
}