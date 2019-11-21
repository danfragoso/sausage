package sausage

import (
	"strings"
)

//Tokenize - Reads source code and tokenizes it, return can be slice with tokens or an error
func Tokenize(source string) interface{} {
	var tokens []*Token
	sourceOriginal := removeImediateWhitespace(source)

	if len(source) == 0 {
		return tokens
	}

	for len(source) > 0 {
		var currentToken string

		if currentToken == "" {
			source = sourceOriginal
			for tokenOrder := 0; tokenOrder < len(tokenTypes); tokenOrder++ {
				tokenIndex := tokenPatterns[tokenTypes[tokenOrder]].FindStringIndex(source)

				if len(tokenIndex) > 0 {

					if tokenIndex[0] == 0 {
						currentToken = getCurrentToken(source, tokenIndex, tokenTypes[tokenOrder])

						tokens = append(tokens, &Token{
							Type:  tokenTypes[tokenOrder],
							Value: currentToken,
							Range: getAbsoluteRange(sourceOriginal, source, tokenIndex),
						})

						source = removeCurrentTokenFromSource(source, tokenIndex)

						tokenOrder = 0
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

func getCurrentToken(source string, tokenIndex []int, tokenType string) string {
	currentToken := source[tokenIndex[0]:tokenIndex[1]]
	if tokenType == "BlockComment" {
		currentToken = strings.Trim(currentToken, "/*")
	} else if tokenType == "LineComment" {
		currentToken = strings.Trim(currentToken, "//")
	}

	return currentToken
}

func removeCurrentTokenFromSource(source string, tokenIndex []int) string {
	return strings.TrimSpace(source[tokenIndex[1]:len(source)])
}
