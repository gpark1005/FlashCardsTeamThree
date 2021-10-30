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
	matching := "matching.json"
	infoOnly := "infoonly.json"
	multiple := "multiplechoice.json"
	tOrF := "torf.json"
	qAndA := "qanda.json"

	files := []string{matching, infoOnly, multiple, tOrF, qAndA}

	for _, v := range files {
		ext := filepath.Ext(v)
		if ext != ".json" {
			log.Fatal("File extension invalid")
		}
	}

	r := repo.NewRepo(matching, infoOnly, multiple, tOrF, qAndA)
	svc := service.NewService(r)
	handler := handlers.NewFlashcardHandler(svc)
	router := handlers.ConfigureRouter(handler)


	svr := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatalln(svr.ListenAndServe())
}
