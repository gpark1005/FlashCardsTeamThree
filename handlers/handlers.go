package handlers

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"io/ioutil"
	"net/http"
	//"github.com/gorilla/mux"
)

type Service interface {
	CreateMatching(f entities.Matching) error
	CreateInfoOnly(f entities.InfoOnly) error
	CreateQAndA(f entities.QAndA) error
	CreateTOrF(f entities.TOrF) error
	CreateMultipleChoice(f entities.MultipleChoice) error
}

type FlashcardHandler struct {
	Svc Service
}

func NewFlashcardHandler(s Service) FlashcardHandler {
	return FlashcardHandler{
		s,
	}

}
var FcType map[string]interface{}

//type TypeCard struct {
//	CType string `json:"type"`
//}


func (fh FlashcardHandler) PostFlashcardHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//getVar := vars["type"]
	//fcType := struct {
	//	Type string `json:"Type"`
	//}{}

	data,err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//tc := TypeCard{}

	err = json.Unmarshal(data, &FcType)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//fmt.Println(tc)

	if cType, ok := FcType["Type"]; ok {
		switch cType{
		case "Matching":
			fcMatching := entities.Matching{}
			err = json.NewDecoder(r.Body).Decode(&fcMatching)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = fh.Svc.CreateMatching(fcMatching)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "InfoOnly":
			fcInfoOnly := entities.InfoOnly{}
			err = json.NewDecoder(r.Body).Decode(&fcInfoOnly)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = fh.Svc.CreateInfoOnly(fcInfoOnly)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "QAndA":
			fcQAndA := entities.QAndA{}
			err = json.NewDecoder(r.Body).Decode(&fcQAndA)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = fh.Svc.CreateQAndA(fcQAndA)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "TorF":
			fcTorF := entities.TOrF{}
			err = json.NewDecoder(r.Body).Decode(&fcTorF)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = fh.Svc.CreateTOrF(fcTorF)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		case "MultipleChoice":
			fcMultipleChoice := entities.MultipleChoice{}
			err = json.NewDecoder(r.Body).Decode(&fcMultipleChoice)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			err = fh.Svc.CreateMultipleChoice(fcMultipleChoice)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}