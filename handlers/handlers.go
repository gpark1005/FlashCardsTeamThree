package handlers

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"github.com/gpark1005/FlashCardsTeamThree/repo"
	"net/http"
	//"github.com/gorilla/mux"
)

type Service interface{
	CreateMatching (f entities.Matching) (*repo.Database, error)
	CreateInfoOnly(f entities.InfoOnly) (*repo.Database, error)
	CreateQAndA(f entities.QAndA) (*repo.Database, error)
	CreateTOrF(f entities.TOrF) (*repo.Database, error)
	CreateMultipleChoice(f entities.MultipleChoice) (*repo.Database, error)
}

type FlashcardHandler struct {
	Svc Service
}

func NewFlashcardHandler(s Service) FlashcardHandler {
	return FlashcardHandler{
		s,
	}

}

func (fh FlashcardHandler) PostFlashcardHandler (w http.ResponseWriter, r *http.Request){
	//vars := mux.Vars(r)
	//getVar := vars["type"]
	fcType := struct {
		Type string `json:"Type"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&fcType)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	switch fcType.Type {
	case "Matching":
		fcMatching := entities.Matching{}
		err = json.NewDecoder(r.Body).Decode(&fcMatching)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		db, err := fh.Svc.CreateMatching(fcMatching)
	case "InfoOnly":
		fcInfoOnly := entities.InfoOnly{}
		err = json.NewDecoder(r.Body).Decode(&fcInfoOnly)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		db, err := fh.Svc.CreateInfoOnly(fcInfoOnly)
	case "QAndA":
		fcQAndA := entities.QAndA{}
		err = json.NewDecoder(r.Body).Decode(&fcQAndA)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		db, err := fh.Svc.CreateQAndA(fcQAndA)
	case "TorF":
		fcTorF := entities.TOrF{}
		err = json.NewDecoder(r.Body).Decode(&fcTorF)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		db, err := fh.Svc.CreateTOrF(fcTorF)
	case "MultipleChoice":
		fcMultipleChoice := entities.MultipleChoice{}
		err = json.NewDecoder(r.Body).Decode(&fcMultipleChoice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		db, err := fh.Svc.CreateMultipleChoice(fcMultipleChoice)
	}
}