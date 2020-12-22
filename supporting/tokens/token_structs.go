package tokens

// basic container token with container id and value, ready for lexing
type ContainerToken struct {
	Id int
	Value string
}

// variable token to hold id of container, variable to update and new value
type VarToken struct {
	Id int
	Variable string
	Value string
}

// function token to hold id of container, function type/name and arguments
type FunctionToken struct {
	Id int
	Function string
	Arguments string
}
