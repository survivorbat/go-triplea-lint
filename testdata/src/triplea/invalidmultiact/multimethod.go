package triplea

type Person struct {
	Name string
}

func (p *Person) SayHello() string {
	return "Hello " + p.Name
}
