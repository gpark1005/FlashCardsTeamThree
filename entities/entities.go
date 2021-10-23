package entities

type Flashcards struct {
	matching       Matching
	infoOnly       InfoOnly
	qAndA          QAndA
	tOrF           TOrF
	multipleChoice MultipleChoice
}

type Matching struct {
	Id       string
	Ctype    string                 `json:"type"`
	Category string                 `json:"category"`
	Question map[string]interface{} `json:"question"`
	Choices  map[string]interface{} `json:"choices"`
	Answers  map[string]interface{} `json:"answers"`
}

type InfoOnly struct {
	Id          string
	Ctype       string `json:"type"`
	Information string `json:"information"`
}

type QAndA struct {
	Id       string
	Ctype    string `json:"type"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type TOrF struct {
	Id       string
	Ctype    string `json:"type"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type MultipleChoice struct {
	Id       string
	Ctype    string                 `json:"type"`
	Category string                 `json:"category"`
	Question string                 `json:"question"`
	Choices  map[string]interface{} `json:"choices"`
	Answer   string                 `json:"answer"`
}
