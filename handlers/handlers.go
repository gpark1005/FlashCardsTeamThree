package handlers

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"github.com/gpark1005/FlashCardsTeamThree/repo"
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
	GetAllFlashcards() (*repo.Database, error)
}

type FlashcardHandler struct {
	Svc Service
}

func NewFlashcardHandler(s Service) FlashcardHandler {
	return FlashcardHandler{
		s,
	}
}

//var FcType map[string]interface{}

func (fh FlashcardHandler) PostFlashcardHandler(w http.ResponseWriter, r *http.Request) {
	fcType := struct {
		Type string `json:"Type"`
	}{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(data, &fcType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch fcType.Type {
	case "Matching":
		fcMatching := entities.Matching{}
		err = json.Unmarshal(data, &fcMatching)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.CreateMatching(fcMatching)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "InfoOnly":
		fcInfoOnly := entities.InfoOnly{}
		err = json.Unmarshal(data, &fcInfoOnly)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.CreateInfoOnly(fcInfoOnly)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "QAndA":
		fcQAndA := entities.QAndA{}
		err = json.Unmarshal(data, &fcQAndA)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.CreateQAndA(fcQAndA)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "TOrF":
		fcTorF := entities.TOrF{}
		err = json.Unmarshal(data, &fcTorF)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.CreateTOrF(fcTorF)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "MultipleChoice":
		fcMultipleChoice := entities.MultipleChoice{}
		err = json.Unmarshal(data, &fcMultipleChoice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.CreateMultipleChoice(fcMultipleChoice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (fh FlashcardHandler) GetAllFlashcards(w http.ResponseWriter, r *http.Request){
	db, err := fh.Svc.GetAllFlashcards()
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fcData, err := json.MarshalIndent(db, "", "	")
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_,err = w.Write(fcData)
}
