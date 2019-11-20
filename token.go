package sausage

import (
	"strings"
)

//Tokenize - Reads source code and tokenizes it, return can be slice with tokens or an error
func Tokenize(source string) interface{} {
	var tokens []*Token
	sourceOriginal := source

	if len(source) == 0 {
		return tokens
	}

	for len(source) > 0 {
		var currentToken string

		if currentToken == "" {
			source = removeImediateWhitespace(source)

			for tokenType, tokenPattern := range tokenPatterns {
				tokenIndex := tokenPattern.FindStringIndex(source)

				if len(tokenIndex) > 0 {
					if tokenIndex[0] == 0 {
						currentToken = getCurrentToken(source, tokenIndex)

						tokens = append(tokens, &Token{
							Type:  tokenType,
							Value: currentToken,
							Range: getAbsoluteRange(sourceOriginal, source, tokenIndex),
						})

						source = removeCurrentTokenFromSource(source, tokenIndex)
						currentToken = ""
					}
				}
			}
		}
	}

	return tokens
}

func removeImediateWhitespace(source string) string {
	return strings.Trim(source, " \n\r")
}

func getAbsoluteRange(originalString string, sourceString string, tokenIndex []int) []int {
	rangeDiff := len(originalString) - len(sourceString)

	var absoluteRange []int
	absoluteRange = append(absoluteRange, tokenIndex[0]+rangeDiff)
	absoluteRange = append(absoluteRange, tokenIndex[1]+rangeDiff)

	return absoluteRange
}

func getCurrentToken(source string, tokenIndex []int) string {
	return source[tokenIndex[0]:tokenIndex[1]]
}

func removeCurrentTokenFromSource(source string, tokenIndex []int) string {
	return strings.TrimSpace(source[tokenIndex[1]:len(source)])
}
