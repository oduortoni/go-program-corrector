package core

type Syntax struct {
	Sourcefile string   `json:"sourcefile"`
	Syntaxfile string   `json:"syntaxfile"`
	Syntax     []string `json:"syntax"`
	Language   string   `json:"language"`
}

type Suggestion struct {
	Key     string
	Suggest []string
}
