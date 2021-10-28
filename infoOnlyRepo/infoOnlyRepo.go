package infoOnlyRepo

import "github.com/gpark1005/FlashCardsTeamThree/entities"

type Database struct {
	InfoOnly       []entities.InfoOnly
}

type InfoOnly struct {
	Filename string
}

func NewInfoOnly(f string) InfoOnly {
	return InfoOnly{
		Filename: f,
	}
}
