package main

import (
	"github.com/gpark1005/FlashCardsTeamThree/handlers"
	"github.com/gpark1005/FlashCardsTeamThree/repo"
	"github.com/gpark1005/FlashCardsTeamThree/service"
	"log"
	"net/http"
	"path/filepath"
)
func main() {
	fn := "flashcards.json"

	ext := filepath.Ext(fn)
	if ext != ".json" {
		log.Fatal("File extension invalid")
	}

	r := repo.NewRepo(fn)
	svc := service.NewService(r)
	handler := handlers.NewFlashcardHandler(svc)
	router := handlers.ConfigureRouter(handler)

	svr := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8080",
	}

	log.Fatalln(svr.ListenAndServe())
}
