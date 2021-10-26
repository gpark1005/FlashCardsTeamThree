package repo

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"github.com/gpark1005/FlashCardsTeamThree/handlers"
	"io/ioutil"
)

type StoreJsonValues struct {
	Types []string
}

func ValidateType(t handlers.FcType) (bool, error) {
	fileType := StoreJsonValues{}
	data, err := ioutil.ReadFile("types.json")
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &fileType)

	for _, v := range fileType.Types {
		if
	}
}


func ValidateMatching(m entities.Matching) (bool, error) {

}


func ValidateInfoOnly(i entities.InfoOnly) (bool, error) {

}


func ValidateQAndA(q entities.QAndA) (bool, error) {

}


func ValidateTOrF(t entities.TOrF) (bool, error) {

}


func ValidateMultipleChoice(mc entities.MultipleChoice) (bool, error) {

}