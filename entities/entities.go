package entities

type Matching struct {
	Type string
	Category string
	Question []string
	Choices map[int]string
	Answers map[string]int
}
