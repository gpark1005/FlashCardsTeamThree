package handlers

import "github.com/gorilla/mux"

func ConfigureRouter(fh FlashcardHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/flashcard", fh.PostFlashcardHandler).Methods("POST")
	r.HandleFunc("/flashcard", fh.GetAllFlashcards).Methods("GET")
	r.HandleFunc("/flashcard/{Id}", fh.GetById).Methods("GET")
	r.HandleFunc("/flashcard/{Id}", fh.UpdateById).Methods("PUT")
	r.HandleFunc("/flashcard/{Id}", fh.DeleteByIdHandler).Methods("DELETE")



	return r
}
