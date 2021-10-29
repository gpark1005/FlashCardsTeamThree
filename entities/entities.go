package entities

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type FcType struct {
	Type string `json:"Type"`
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