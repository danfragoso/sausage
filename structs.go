package sausage

/*Sausage - VM Struct*/
type Sausage struct {
	Tree *AST
}

/*AST - VM AST Struct*/
type AST struct {
	String string
}

//GetJSON - Returns the AST in JSON format
func (ast AST) GetJSON() string {
	return ast.String
}
