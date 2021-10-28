package qandaRepo

import "github.com/gpark1005/FlashCardsTeamThree/entities"

type Database struct {
	QAndA      []entities.QAndA
}

type QAndA struct {
	Filename string
}

func NewQAndA(f string) QAndA {
	return QAndA{
		Filename: f,
	}
}
