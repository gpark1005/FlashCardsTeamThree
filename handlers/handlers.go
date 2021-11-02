package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
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
	GetAllFlashcards() (*repo.Flashcards, error)
	GetMatchingById(id string) (*entities.Matching, error)
	GetInfoOnlyById(id string) (*entities.InfoOnly, error)
	GetQAndAById(id string) (*entities.QAndA, error)
	GetTOrFById(id string) (*entities.TOrF, error)
	GetMultipleChoiceById(id string) (*entities.MultipleChoice, error)
	UpdateMatchingById(id string, m entities.Matching) error
	UpdateInfoById(id string, io entities.InfoOnly) error
	UpdateQandAById(id string, qa entities.QAndA) error
	UpdateMultipleChoiceById(id string, mc entities.MultipleChoice) error
	UpdateTorFById(id string, tf entities.TOrF) error
	DeleteMatchingById(id string) error
	DeleteInfoOnlyById(id string)error
	DeleteTOrFById (id string) error
	DeleteQAndAById (id string) error
	DeleteMCById (id string) error
}

type FlashcardHandler struct {
	Svc Service
}

func NewFlashcardHandler(s Service) FlashcardHandler {
	return FlashcardHandler{
		s,
	}
}

func (fh FlashcardHandler) PostFlashcardHandler(w http.ResponseWriter, r *http.Request) {
	cType := entities.FcType{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(data, &cType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valType, err := entities.ValidateType(cType)
	if err != nil {
		switch err.Error(){
		case "type not valid":
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	if valType == true {
		switch cType.Type {
		case "Matching":
			fcMatching := entities.Matching{}
			err = json.Unmarshal(data, &fcMatching)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			err = fcMatching.ValidateMatching()
			if err != nil {
				switch err.Error(){
				case "category required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "answer key required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "choices required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "questions required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
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
			err = fcInfoOnly.ValidateInfoOnly()
			if err != nil {
				switch err.Error(){
				case "category required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "information required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
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
			err = fcQAndA.ValidateQAndA()
			if err != nil{
				switch err.Error() {
				case "category required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "question(s) required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "answer required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
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
			err = fcTorF.ValidateTOrF()
			if err != nil{
				switch err.Error() {
				case "category required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "question(s) required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "answer required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
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
			err = fcMultipleChoice.ValidateMultipleChoice()
			if err != nil{
				switch err.Error() {
				case "category required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "question(s) required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "answer required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "choices required":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}
			err = fh.Svc.CreateMultipleChoice(fcMultipleChoice)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

func (fh FlashcardHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars ["Id"]

	mFc, err := fh.Svc.GetMatchingById(id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	ioFc, err := fh.Svc.GetInfoOnlyById(id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	qaFc, err := fh.Svc.GetQAndAById(id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	tfFc, err := fh.Svc.GetTOrFById(id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	mcFc, err := fh.Svc.GetMultipleChoiceById(id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	switch {
	case mFc != nil:
		mData, err := json.MarshalIndent(mFc, "", "	")
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, err = w.Write(mData)
	case ioFc != nil:
		ioData, err := json.MarshalIndent(ioFc, "", "	")
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, err = w.Write(ioData)
	case qaFc != nil:
		qaData, err := json.MarshalIndent(qaFc, "", "	")
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, err = w.Write(qaData)
	case tfFc != nil:
		tfData, err := json.MarshalIndent(tfFc, "", "	")
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, err = w.Write(tfData)
	case mcFc != nil:
		mcData, err := json.MarshalIndent(mcFc, "", "	")
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, err = w.Write(mcData)
	default:
		err = errors.New("flashcard not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}

func (fh FlashcardHandler) UpdateById(w http.ResponseWriter, r *http.Request){
	fhId := mux.Vars(r)
	id := fhId["Id"]
	cType := entities.FcType{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
	err = json.Unmarshal(data, &cType)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
	switch cType.Type {
	case "Matching":
		fcMatching := entities.Matching{}
		err = json.Unmarshal(data, &fcMatching)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.UpdateMatchingById(id, fcMatching)
	case "InfoOnly":
		fcInfoOnly := entities.InfoOnly{}
		err = json.Unmarshal(data, &fcInfoOnly)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.UpdateInfoById(id, fcInfoOnly)
	case "QAndA":
		fcQAndA := entities.QAndA{}
		err = json.Unmarshal(data, &fcQAndA)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.UpdateQandAById(id, fcQAndA)
	case "TOrF":
		fcTorF := entities.TOrF{}
		err = json.Unmarshal(data, &fcTorF)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.UpdateTorFById(id, fcTorF)
	case "MultipleChoice":
		fcMultipleChoice := entities.MultipleChoice{}
		err = json.Unmarshal(data, &fcMultipleChoice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = fh.Svc.UpdateMultipleChoiceById(id, fcMultipleChoice)
	default:
		err = errors.New("flashcard not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (fh FlashcardHandler) DeleteByIdHandler (w http.ResponseWriter, r *http.Request) {
	fhId := mux.Vars(r)
	id := fhId["Id"]

	err := fh.Svc.DeleteMatchingById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	err = fh.Svc.DeleteInfoOnlyById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	err = fh.Svc.DeleteTOrFById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	err = fh.Svc.DeleteQAndAById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	err = fh.Svc.DeleteMCById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	http.Error(w, err.Error() , http.StatusNoContent)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
