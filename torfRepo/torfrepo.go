package torfRepo

import "github.com/gpark1005/FlashCardsTeamThree/entities"

type Database struct {
	TOrF      []entities.TOrF
}

type TOrF struct {
	Filename string
}

func NewTOrF(f string) TOrF {
	return TOrF{
		Filename: f,
	}
}
