package sausage

//Tokenize - Reads source code and tokenizes it, return can be slice with tokens or an error
func Tokenize(source string) interface{} {
	var tokens []*Token

	if len(source) == 0 {
		return tokens
	}

	return nil
}
