package tokens

// basic container token with container id and value, ready for lexing
type ContainerToken struct {
	Id int
	Value string
}
