package sausage

/*Token - Token format*/
type Token struct {
	Type  string   `json:"type"`
	Value string   `json:"value"`
	Range [2]int32 `json:"range"`
}

/*TokenizerFailure - Tokenizer failure format*/
type TokenizerFailure struct {
	Tokenize   bool   `json:"tokenize"`
	Index      int32  `json:"index"`
	LineNumber int32  `json:"lineNumber"`
	Column     int32  `json:"column"`
	Message    string `json:"message"`
}
