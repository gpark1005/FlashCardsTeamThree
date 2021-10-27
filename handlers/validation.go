package handlers

import (
	"encoding/json"
	"errors"
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

//
//func ValidateMatching(m entities.Matching) (bool, error) {
//
//}
//
//
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