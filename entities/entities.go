package entities

type Flashcards struct {
	Matching Matching
	InfoOnly InfoOnly
	QAndA QAndA
	TOrF TorF
	MultipleChoice MultipleChoice
}

type Question struct {
	Questions []string
}

type Choices struct {
	Choices map[int]string
}

type Answers struct {
	Answers map[string]int
}

type Matching struct {
	Type string
	Category string
	Question Question
	Choices Choices
	Answers Answers
}

type InfoOnly struct {
	Id string
	Type string
	Information string
}

type QAndA struct {
	Type string
	Category string
	Question string
	Answer string
}

type TorF struct {
	Type string
	Category string
	Question string
	Answer string
}

type MultipleChoice struct {
	Type string
	Category string
	Question string
	Choices Choices
	Answer string
}