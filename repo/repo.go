package repo

import (
	"github.com/gpark1005/FlashCardsTeamThree/entities"

)

type FlashCards struct {
	Matching entities.Matching
	InfoOnly entities.InfoOnly
	QAndA entities.QAndA
	TOrF entities.TorF
	MultipleChoice entities.MultipleChoice

}