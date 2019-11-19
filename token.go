package sausage

import (
	"fmt"
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
			for tokenType, tokenPattern := range tokenPatterns {
				tokenIndex := tokenPattern.FindStringIndex(source)

				if len(tokenIndex) > 0 {
					if tokenIndex[0] == 0 {
						rangeDiff := len(sourceOriginal) - len(source)

						var absoluteRange []int

						currentToken = source[tokenIndex[0]:tokenIndex[1]]

						absoluteRange = append(absoluteRange, tokenIndex[0]+rangeDiff)
						absoluteRange = append(absoluteRange, tokenIndex[1]+rangeDiff)

						tokens = append(tokens, &Token{
							Type:  tokenType,
							Value: currentToken,
							Range: absoluteRange,
						})

						source = strings.TrimSpace(source[tokenIndex[1]:len(source)])

						currentToken = ""
					}
				}
			}
		}
	}

	for _, v := range tokens {
		fmt.Println(v)
	}
	return tokens
}
