package matchingRepo

import "github.com/gpark1005/FlashCardsTeamThree/entities"

type Database struct {
	Matching       []entities.Matching
}

type Matching struct {
	Filename string
}

func NewMatching(f string) Matching {
	return Matching{
		Filename: f,
	}
}
