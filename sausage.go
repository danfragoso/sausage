package sausage

//New - Creates and returns a new sausage vm
func New() *Sausage {
	AST := AST{}

	return &Sausage{
		Tree: &AST,
	}
}
