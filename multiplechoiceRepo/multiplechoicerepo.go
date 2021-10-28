package multiplechoiceRepo

import "github.com/gpark1005/FlashCardsTeamThree/entities"

type Database struct {
	MultipleChoice       []entities.MultipleChoice
}

type MultipleChoice struct {
	Filename string
}

func NewMultipleChoice(f string) MultipleChoice {
	return MultipleChoice{
		Filename: f,
	}
}
