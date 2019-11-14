package sausage

import (
	"errors"
	"io/ioutil"
)

//LoadTree - Loads a AST String on the VM
func (vm Sausage) LoadTree(tree string) error {
	if len(tree) == 0 {
		return errors.New("Tree string is empty")
	}

	vm.Tree.String = tree
	return nil
}

//LoadTreeFile - Loads a AST file on the VM
func (vm Sausage) LoadTreeFile(path string) error {
	if len(path) == 0 {
		return errors.New("The path is empty")
	}

	buffer, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	vm.Tree.String = string(buffer)
	return nil
}
