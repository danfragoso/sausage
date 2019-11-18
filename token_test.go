package sausage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	files, err := ioutil.ReadDir("tests/token")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(files); i += 2 {
		fileName := files[i].Name()
		testName := strings.Split(fileName, ".")[0]
		tokenizerErrorFlag := false
		result := "{}"

		sourceBuffer, sourceErr := ioutil.ReadFile("tests/token/" + testName + ".js")
		if sourceErr != nil {
			errorString := "Error loading test" + testName
			t.Error(errorString)
		}

		source := string(sourceBuffer)

		tokensBuffer, tokensErr := ioutil.ReadFile("tests/token/" + testName + ".tokens.json")
		result = string(tokensBuffer)

		if tokensErr != nil {
			failureBuffer, failureErr := ioutil.ReadFile("tests/token/" + testName + ".failure.json")
			result = string(failureBuffer)
			tokenizerErrorFlag = true

			if failureErr != nil {
				errorString := "Error loading test: " + testName
				t.Error(errorString)
			}
		}

		tokenizeResult := Tokenize(source)

		if tokenizerErrorFlag {
			var expected *TokenizerFailure

			err := json.Unmarshal([]byte(result), &expected)
			if err != nil {
				t.Error(err)
			}

			testPass := assert.Equal(t, expected, tokenizeResult, "Error testing: "+testName)
			if testPass {
				fmt.Println("└───── PASS - " + testName)
			}
		} else {
			var expected []*Token
			err := json.Unmarshal([]byte(result), &expected)
			if err != nil {
				t.Error(err)
			}

			if reflect.TypeOf(tokenizeResult) == reflect.TypeOf(expected) {
				testPass := assert.ElementsMatch(t, expected, tokenizeResult, "Error testing "+testName)
				if testPass {
					fmt.Println("└───── PASS - " + testName)
				}
			} else {
				testPass := assert.Equal(t, expected, tokenizeResult, "Error testing "+testName+": Types should be equal")
				if testPass {
					fmt.Println("└───── PASS - " + testName)
				}
			}
		}
	}
}
