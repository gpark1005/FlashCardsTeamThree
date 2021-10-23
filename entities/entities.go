package entities

type Flashcards struct {
	Matching       Matching
	InfoOnly       InfoOnly
	QAndA          QAndA
	TOrF           TOrF
	MultipleChoice MultipleChoice
}

type Matching struct {
	Id       string
	Type     string
	Category string
	Question map[string]interface{}
	Choices  map[string]interface{}
	Answers  map[string]interface{}
}

type InfoOnly struct {
	Id          string
	Type        string
	Information string
}

type QAndA struct {
	Id       string
	Type     string
	Category string
	Question string
	Answer   string
}

type TOrF struct {
	Id       string
	Type     string
	Category string
	Question string
	Answer   string
}

type MultipleChoice struct {
	Id       string
	Type     string
	Category string
	Question string
	Choices  map[string]interface{}
	Answer   string
}
