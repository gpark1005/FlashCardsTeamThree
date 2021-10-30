package handlers

import (
	"encoding/json"
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
	GetAllFlashcards() (*repo.Database, error)
	//GetById(id string) (map[string]interface{}, error)
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
//var FcType map[string]interface{}

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

//func (fh FlashcardHandler) GetById(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars ["Id"]
//
//	fc, err := fh.Svc.GetById(id)
//	if err != nil{
//		switch err.Error() {
//		case "flashcard does not exist":
//			http.Error(w, err.Error(), http.StatusNotFound)
//		}
//	}
//
//	if err != nil{
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//
//	fcData, err := json.MarshalIndent(fc, "", "	")
//	if err != nil{
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusAccepted)
//	_,err = w.Write(fcData)
//}

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}



