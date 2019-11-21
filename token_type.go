package sausage

import "regexp"

/*Token - Token format*/
type Token struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Range []int  `json:"range"`
}

/*TokenizerFailure - Tokenizer failure format*/
type TokenizerFailure struct {
	Tokenize   bool   `json:"tokenize"`
	Index      int32  `json:"index"`
	LineNumber int32  `json:"lineNumber"`
	Column     int32  `json:"column"`
	Message    string `json:"message"`
}

var tokenPatterns = map[string]*regexp.Regexp{
	"Keyword":           regexp.MustCompile(`^abstract|arguments|await|boolean|break|byte|case|catch|char|class|const|continue|debugger|default|delete|do|double|else|enum|eval|export|extends|false|final|finally|float|for|function|goto|if|implements|import|in|instanceof|int|interface|let|long|native|new|null|package|private|protected|public|return|short|static|super|switch|synchronized|this|throw|throws|transient|true|try|typeof|var|void|volatile|while|with|yield`),
	"Null":              regexp.MustCompile(`^null`),
	"Boolean":           regexp.MustCompile(`^true|false`),
	"Numeric":           regexp.MustCompile(`^\d+`),
	"String":            regexp.MustCompile(`^["|'].*["|']`),
	"RegularExpression": regexp.MustCompile(`^\/[^*]\S*\/y?m?g?i?u?s?`),
	"LineComment":       regexp.MustCompile(`^\/\/.*\n?`),
	"BlockComment":      regexp.MustCompile(`^\/\*.*\*\/`),
	"Punctuator":        regexp.MustCompile(`^{|\(|\)|\[|\]|\.{3}|\.|;|,|<<=|>>>=|>>=|>{3}|>{2}|<{2}|<=|>=|=>|<|>|===|!==|==|!=|\+=|\+{2}|\+|\-=|-{2}|\-|\*{2}=|\*{2}|\*=|\*|%=|%|&{2}|&=|&|\|{2}|\|=|\||~|\?|:|\^=|\^|!|/=|=|}|/`),
	"Identifier":        regexp.MustCompile(`^[a-zA-z|$]+`),
	"Template":          regexp.MustCompile(`^\x60.*\x60`),
}

var tokenTypes = []string{
	"Keyword",
	"Null",
	"Boolean",
	"Numeric",
	"String",
	"RegularExpression",
	"LineComment",
	"BlockComment",
	"Punctuator",
	"Identifier",
	"Template",
}
