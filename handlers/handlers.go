package handlers

import (
	""
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

}