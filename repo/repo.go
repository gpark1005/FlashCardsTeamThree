package repo

import (
	"encoding/json"
	"github.com/gpark1005/FlashCardsTeamThree/entities"
	//"github.com/gpark1005/FlashCardsTeamThree/infoOnlyRepo"
	//"github.com/gpark1005/FlashCardsTeamThree/matchingRepo"
	//"github.com/gpark1005/FlashCardsTeamThree/multiplechoiceRepo"
	//"github.com/gpark1005/FlashCardsTeamThree/qandaRepo"
	//"github.com/gpark1005/FlashCardsTeamThree/torfRepo"
	"io/ioutil"
)

type Database struct {
	Flashcards []interface{}
}

type flashcards struct {
	Flashcards []map[string]interface{}
}

type Repo struct {
	MatchingFile string
	InfoOnlyFile string
	QAndAFile string
	MultipleChoiceFile string
	TrueOrFalseFile string
}

func NewRepo(m string, io string, mc string, tf string, qa string) Repo {
	return Repo{
		MatchingFile: m,
		InfoOnlyFile: io,
		QAndAFile: qa,
		MultipleChoiceFile: mc,
		TrueOrFalseFile: tf,
	}
}

func (r Repo) CreateMatching(f entities.Matching) error {
	fcSlice := entities.MatchingDatabase{}

	file, err := ioutil.ReadFile(r.MatchingFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.Matching = append(fcSlice.Matching, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.MatchingFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateInfoOnly(f entities.InfoOnly) error {
	fcSlice := entities.InfoOnlyDatabase{}

	file, err := ioutil.ReadFile(r.InfoOnlyFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.InfoOnly = append(fcSlice.InfoOnly, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.InfoOnlyFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateQAndA(f entities.QAndA) error {
	fcSlice := entities.QAndADatabase{}

	file, err := ioutil.ReadFile(r.QAndAFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.QAndA = append(fcSlice.QAndA, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.QAndAFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateTOrF(f entities.TOrF) error {
	fcSlice := entities.TOrFDatabase{}

	file, err := ioutil.ReadFile(r.TrueOrFalseFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.TOrF = append(fcSlice.TOrF, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.TrueOrFalseFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CreateMultipleChoice(f entities.MultipleChoice) error {
	fcSlice := entities.MCDatabase{}

	file, err := ioutil.ReadFile(r.MultipleChoiceFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &fcSlice)
	if err != nil {
		return err
	}

	fcSlice.MultipleChoice = append(fcSlice.MultipleChoice, f)

	fcBytes, err := json.MarshalIndent(fcSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.MultipleChoiceFile, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}


func (r Repo) GetAllFlashcards() (*Database, error) {
	fcSlice := Database{}
	files := []string{r.MatchingFile, r.MultipleChoiceFile, r.TrueOrFalseFile, r.QAndAFile, r.InfoOnlyFile}

	for _, v := range files {
		file, err := ioutil.ReadFile(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(file, &fcSlice)
		if err != nil {
			return nil, err
		}
	}
	return &fcSlice, nil
}

//func (r Repo) GetById(id string) (map[string]interface{}, error) {
//
//	db := flashcards{}
//
//	file, err := ioutil.ReadFile()
//	if err != nil {
//		return nil, err
//	}
//
//	err = json.Unmarshal(file, &db)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, v := range db.Flashcards{
//		if idCheck, ok := v["Id"]; ok{
//			if idCheck == id{
//				return v, nil
//			}
//		}
//	}
//	return nil, errors.New("flashcard does not exist")
//}

func (r Repo) UpdateMatchingById(id string, m entities.Matching) error {
	file, err := ioutil.ReadFile(r.MatchingFile)
	if err != nil{
		return err
	}

	md := entities.MatchingDatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.Matching{
		if v.Id == id{
			m.Id = id
			md.Matching[i] = m
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.MatchingFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) UpdateInfoById(id string, io entities.InfoOnly) error{
	file, err := ioutil.ReadFile(r.InfoOnlyFile)
	if err != nil{
		return err
	}

	md := entities.InfoOnlyDatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.InfoOnly{
		if v.Id == id{
			io.Id = id
			md.InfoOnly[i] = io
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.InfoOnlyFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) UpdateQandAById(id string, qa entities.QAndA) error{
	file, err := ioutil.ReadFile(r.QAndAFile)
	if err != nil{
		return err
	}

	md := entities.QAndADatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.QAndA{
		if v.Id == id{
			qa.Id = id
			md.QAndA[i] = qa
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.QAndAFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) UpdateMultipleChoiceById(id string, mc entities.MultipleChoice) error{
	file, err := ioutil.ReadFile(r.MultipleChoiceFile)
	if err != nil{
		return err
	}

	md := entities.MCDatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.MultipleChoice{
		if v.Id == id{
			mc.Id = id
			md.MultipleChoice[i] = mc
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.MultipleChoiceFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) UpdateTorFById(id string, tf entities.TOrF) error {
	file, err := ioutil.ReadFile(r.TrueOrFalseFile)
	if err != nil{
		return err
	}

	md := entities.TOrFDatabase{}
	err= json.Unmarshal(file, &md)

	for i, v := range md.TOrF{
		if v.Id == id{
			tf.Id = id
			md.TOrF[i] = tf
		}
	}
	marshal, err := json.MarshalIndent(md, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.TrueOrFalseFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) DeleteMatchingById (id string) error{
	m := entities.MatchingDatabase{}

	file, err := ioutil.ReadFile(r.MatchingFile)
	if err != nil{
		return err
	}

	err = json.Unmarshal(file, &m)
	if err != nil{
		return err
	}

	for i, v := range m.Matching{
		if v.Id == id{
			m.Matching = append(m.Matching[:i], m.Matching[i+1:]...)
		}
	}

	marshal, err := json.MarshalIndent(m, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.MatchingFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) DeleteInfoOnlyById (id string) error{
	i := entities.InfoOnlyDatabase{}

	file, err := ioutil.ReadFile(r.InfoOnlyFile)
	if err != nil{
		return err
	}

	err = json.Unmarshal(file, &i)
	if err != nil{
		return err
	}

	for index, v := range i.InfoOnly{
		if v.Id == id{
			i.InfoOnly = append(i.InfoOnly[:index], i.InfoOnly[index+1:]...)
		}
	}

	marshal, err := json.MarshalIndent(i, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.InfoOnlyFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) DeleteQAndAById (id string) error{
	q := entities.QAndADatabase{}

	file, err := ioutil.ReadFile(r.QAndAFile)
	if err != nil{
		return err
	}

	err = json.Unmarshal(file, &q)
	if err != nil{
		return err
	}

	for i, v := range q.QAndA{
		if v.Id == id{
			q.QAndA = append(q.QAndA[:i], q.QAndA[i+1:]...)
		}
	}

	marshal, err := json.MarshalIndent(q, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.QAndAFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) DeleteTOrFById (id string) error{
	t := entities.TOrFDatabase{}

	file, err := ioutil.ReadFile(r.TrueOrFalseFile)
	if err != nil{
		return err
	}

	err = json.Unmarshal(file, &t)
	if err != nil{
		return err
	}

	for i, v := range t.TOrF{
		if v.Id == id{
			t.TOrF = append(t.TOrF[:i], t.TOrF[i+1:]...)
		}
	}

	marshal, err := json.MarshalIndent(t, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.TrueOrFalseFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}

func (r Repo) DeleteMCById (id string) error{
	mc := entities.MCDatabase{}

	file, err := ioutil.ReadFile(r.MultipleChoiceFile)
	if err != nil{
		return err
	}

	err = json.Unmarshal(file, &mc)
	if err != nil{
		return err
	}

	for i, v := range mc.MultipleChoice{
		if v.Id == id{
			mc.MultipleChoice = append(mc.MultipleChoice[:i], mc.MultipleChoice[i+1:]...)
		}
	}

	marshal, err := json.MarshalIndent(mc, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.MultipleChoiceFile, marshal, 0644)
	if err != nil{
		return err
	}
	return nil
}