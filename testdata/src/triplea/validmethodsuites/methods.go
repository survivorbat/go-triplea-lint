package validmethods

type MyService struct {
	Name string
}

func (m *MyService) SayHello() string {
	return "Hello " + m.Name
}

func (m *MyService) SayGoodbye() string {
	return "Goodbye " + m.Name
}

func (m *MyService) SayGoodMorning() string {
	return "Good morning " + m.Name
}

func (m *MyService) SayGoodAfternoon() string {
	return "Good afternoon " + m.Name
}
