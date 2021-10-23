package entities

type Flashcards struct {
	matching       Matching
	infoOnly       InfoOnly
	qAndA          QAndA
	tOrF           TOrF
	multipleChoice MultipleChoice
}

type Matching struct {
	id       string
	ctype     string`json:"type"`
	category string `json:"category"`
	question map[string]interface{}`json:"question"`
	choices  map[string]interface{} `json:"choices"`
	answers  map[string]interface{}`json:"answers"`
}

type InfoOnly struct {
	id          string
	ctype        string `json:"type"`
	information string `json:"information"`
}

type QAndA struct {
	id       string
	ctype     string `json:"type"`
	category string `json:"category"`
	question string`json:"question"`
	answer   string`json:"answer"`
}

type TOrF struct {
	id       string
	ctype     string `json:"type"`
	category string`json:"category"`
	question string`json:"question"`
	answer   string`json:"answer"`
}

type MultipleChoice struct {
	id       string
	ctype     string `json:"type"`
	category string`json:"category"`
	question string`json:"question"`
	choices  map[string]interface{} `json:"choices"`
	answer   string `json:"answer"`
}
