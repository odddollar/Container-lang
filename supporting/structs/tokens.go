package structs

// basic container token with container id and value, ready for lexing
type ContainerToken struct {
	Id int
	Value string
}

// common token struct, acts as the parent for variable and function token structs
type Token struct {
	Id int
	VarToken VarToken
	FunctionToken FunctionToken
	Block []Token
}

// variable token to hold id of container, variable to update and new value
type VarToken struct {
	Variable string
	Value string
}

// function token to hold id of container, function type/name and arguments
type FunctionToken struct {
	Function string
	Arguments string
}
