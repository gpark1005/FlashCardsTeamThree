package handlers

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	"github.com/gpark1005/FlashCardsTeamThree/repo"
	"net/http"
	//"github.com/gorilla/mux"
)

type Service interface{
	CreateFlashcard (f entities.Flashcards) (*repo.Database, error)
}

type FlashcardHandler struct {
	Svc Service
}

func NewFlashcardHandler(s Service) FlashcardHandler {
	return FlashcardHandler{
		s,
	}

}

func (fh FlashcardHandler) PostFlashCardHandler (w http.ResponseWriter, r *http.Request){
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
		flashcards := entities.Flashcards{}
		fcMatching := flashcards.Matching
		err = json.NewDecoder(r.Body).Decode(&fcMatching)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		//fh.Svc.CreateFlashcard(fcMatching)
	case "InfoOnly":
		fcInfoOnly := entities.InfoOnly{}
		err = json.NewDecoder(r.Body).Decode(&fcInfoOnly)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "QAndA":
		fcQAndA := entities.QAndA{}
		err = json.NewDecoder(r.Body).Decode(&fcQAndA)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "TorF":
		fcTorF := entities.TorF{}
		err = json.NewDecoder(r.Body).Decode(&fcTorF)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "MultipleChoice":
		fcMultipleChoice := entities.MultipleChoice{}
		err = json.NewDecoder(r.Body).Decode(&fcMultipleChoice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}