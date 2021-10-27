package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"io/ioutil"
)

type StoreJsonValues struct {
	Types []string
}

func ValidateType(t FcType) (bool, error) {
	fileType := StoreJsonValues{}
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


func ValidateMatching(m entities.Matching) error {
	if m.Answers == nil{
		return errors.New("no answers provided")
	}else if m.Choices == nil{
		return errors.New("no choices provided")
	}else if m.Question == nil{
		return errors.New("no questions provided")
	}
	return nil
}


//func ValidateInfoOnly(i entities.InfoOnly) (bool, error) {
//
//}
//
//
//func ValidateQAndA(q entities.QAndA) (bool, error) {
//
//}
//
//
//func ValidateTOrF(t entities.TOrF) (bool, error) {
//
//}
//
//
//func ValidateMultipleChoice(mc entities.MultipleChoice) (bool, error) {
//
//}