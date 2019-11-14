package sausage

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTree(t *testing.T) {
	buffer, err := ioutil.ReadFile("tests/assets/_1_ast.json")
	if err != nil {
		log.Panic(err)
	}

	treeString := string(buffer)
	vm := New()

	vmErr := vm.LoadTree(treeString)
	assert.Equal(t, nil, vmErr, "Error loading the tree")

	astJSON := vm.Tree.String
	assert.Equal(t, treeString, astJSON, "Tree JSON different from source file")
}
func TestLoadTreeFile(t *testing.T) {
	filePath := "tests/assets/_1_ast.json"

	vm := New()

	vmErr := vm.LoadTreeFile(filePath)
	assert.Equal(t, nil, vmErr, "Error loading the tree file")

	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	treeString := string(buffer)
	astJSON := vm.Tree.String

	assert.Equal(t, treeString, astJSON, "Tree JSON different from source file")
}
