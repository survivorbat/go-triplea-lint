package triplea

// For valid tests

func SayHello(name string) string {
	return "Hello " + name
}

func SayGoodbye(name string) string {
	return "Goodbye " + name
}

func SayGoodMorning(name string) string {
	return "Good morning " + name
}

func SayGoodLateMorning(name string) string {
	return "Good late morning " + name
}

func SayGoodEvening(name string) string {
	return "Good evening " + name
}

func SayGoodAfternoon(name string) string {
	return "Good afternoon " + name
}

func SayGoodDay(name string) string {
	return "Good day " + name
}

func SayGoodNight(name string) string {
	return "Good night " + name
}

// For invalid tests

func WaveHello(name string) string {
	return "👋 Hello " + name
}

func WaveGoodbye(name string) string {
	return "👋 Goodbye " + name
}

func WaveGoodMorning(name string) string {
	return "👋 Good morning " + name
}

func WaveGoodEvening(name string) string {
	return "👋 Good evening " + name
}

func WaveGoodAfternoon(name string) string {
	return "👋 Good afternoon " + name
}

func WaveGoodDay(name string) string {
	return "👋 Good day " + name
}

func WaveGoodNight(name string) string {
	return "👋 Good night " + name
}

// Types for expression tests

type Person struct {
	Name string
	Age  int
}

type Address struct {
	Street string
	City   string
}

type Company struct {
	Name    string
	Address Address
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func NewCompany(name string) *Company {
	return &Company{Name: name}
}

func CreateAddress(street, city string) Address {
	return Address{Street: street, City: city}
}

func GetPersonAge(p *Person) int {
	return p.Age
}

func (p *Person) Greet() string {
	return "Hello, I'm " + p.Name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

// For suite tests

type UserService struct {
	users map[string]*Person
}

func NewUserService() *UserService {
	return &UserService{users: make(map[string]*Person)}
}

func (s *UserService) AddUser(p *Person) {
	s.users[p.Name] = p
}

func (s *UserService) GetUser(name string) (*Person, bool) {
	user, ok := s.users[name]
	return user, ok
}

func (s *UserService) DeleteUser(name string) {
	delete(s.users, name)
}
