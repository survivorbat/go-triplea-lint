package triplea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPerson(t *testing.T) {
	person := NewPerson("Josh", 30)      // want `// Act statement expected`
	assert.Equal(t, "Josh", person.Name) // want `// Assert statement expected`
}

func TestNewPersonValid(t *testing.T) {
	// Act
	person := NewPerson("Josh", 30)

	// Assert
	assert.Equal(t, "Josh", person.Name)
}

func TestPersonStructInit(t *testing.T) {
	person := Person{Name: "Josh", Age: 30} // want `// Act statement expected`
	assert.Equal(t, "Josh", person.Name)    // want `// Assert statement expected`
}

func TestPersonStructInitValid(t *testing.T) {
	// Act
	person := Person{Name: "Josh", Age: 30}

	// Assert
	assert.Equal(t, "Josh", person.Name)
}

func TestPersonPointerInit(t *testing.T) {
	person := &Person{Name: "Josh", Age: 30} // want `// Act statement expected`
	assert.Equal(t, "Josh", person.Name)     // want `// Assert statement expected`
}

func TestPersonPointerInitValid(t *testing.T) {
	// Act
	person := &Person{Name: "Josh", Age: 30}

	// Assert
	assert.Equal(t, "Josh", person.Name)
}

func TestGetPersonAge(t *testing.T) {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}

	age := GetPersonAge(person) // want `// Act statement expected`
	assert.Equal(t, 30, age)    // want `// Assert statement expected`
}

func TestGetPersonAgeValid(t *testing.T) {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}

	// Act
	age := GetPersonAge(person)

	// Assert
	assert.Equal(t, 30, age)
}

func TestPersonGreet(t *testing.T) {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}

	greeting := person.Greet()                   // want `// Act statement expected`
	assert.Equal(t, "Hello, I'm Josh", greeting) // want `// Assert statement expected`
}

func TestPersonGreetValid(t *testing.T) {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}

	// Act
	greeting := person.Greet()

	// Assert
	assert.Equal(t, "Hello, I'm Josh", greeting)
}

func TestPersonFieldAccess(t *testing.T) {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}

	name := person.Name           // want `// Act statement expected`
	assert.Equal(t, "Josh", name) // want `// Assert statement expected`
}

func TestPersonFieldAccessValid(t *testing.T) {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}

	// Act
	name := person.Name

	// Assert
	assert.Equal(t, "Josh", name)
}

func TestCreateAddress(t *testing.T) {
	address := CreateAddress("Main St", "NYC") // want `// Act statement expected`
	assert.Equal(t, "Main St", address.Street) // want `// Assert statement expected`
}

func TestCreateAddressValid(t *testing.T) {
	// Act
	address := CreateAddress("Main St", "NYC")

	// Assert
	assert.Equal(t, "Main St", address.Street)
}

func TestNestedFieldAccess(t *testing.T) {
	// Arrange
	company := &Company{
		Name:    "Acme",
		Address: Address{Street: "Main St", City: "NYC"},
	}

	city := company.Address.City // want `// Act statement expected`
	assert.Equal(t, "NYC", city) // want `// Assert statement expected`
}

func TestNestedFieldAccessValid(t *testing.T) {
	// Arrange
	company := &Company{
		Name:    "Acme",
		Address: Address{Street: "Main St", City: "NYC"},
	}

	// Act
	city := company.Address.City

	// Assert
	assert.Equal(t, "NYC", city)
}

func TestSliceInit(t *testing.T) {
	slice := []string{"a", "b", "c"} // want `// Act statement expected`
	assert.Len(t, slice, 3)          // want `// Assert statement expected`
}

func TestSliceInitValid(t *testing.T) {
	// Act
	slice := []string{"a", "b", "c"}

	// Assert
	assert.Len(t, slice, 3)
}

func TestMapInit(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2} // want `// Act statement expected`
	assert.Len(t, m, 2)                 // want `// Assert statement expected`
}

func TestMapInitValid(t *testing.T) {
	// Act
	m := map[string]int{"a": 1, "b": 2}

	// Assert
	assert.Len(t, m, 2)
}

func TestSliceIndexAccess(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	item := slice[0]           // want `// Act statement expected`
	assert.Equal(t, "a", item) // want `// Assert statement expected`
}

func TestSliceIndexAccessValid(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	// Act
	item := slice[0]

	// Assert
	assert.Equal(t, "a", item)
}

func TestMapKeyAccess(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1, "b": 2}

	value := m["a"]           // want `// Act statement expected`
	assert.Equal(t, 1, value) // want `// Assert statement expected`
}

func TestMapKeyAccessValid(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1, "b": 2}

	// Act
	value := m["a"]

	// Assert
	assert.Equal(t, 1, value)
}

func TestTypeAssertion(t *testing.T) {
	// Arrange
	var i interface{} = "hello"

	str := i.(string)             // want `// Act statement expected`
	assert.Equal(t, "hello", str) // want `// Assert statement expected`
}

func TestTypeAssertionValid(t *testing.T) {
	// Arrange
	var i interface{} = "hello"

	// Act
	str := i.(string)

	// Assert
	assert.Equal(t, "hello", str)
}

func TestTypeAssertionWithOk(t *testing.T) {
	// Arrange
	var i interface{} = "hello"

	str, ok := i.(string) // want `// Act statement expected`
	assert.True(t, ok)    // want `// Assert statement expected`
	assert.Equal(t, "hello", str)
}

func TestTypeAssertionWithOkValid(t *testing.T) {
	// Arrange
	var i interface{} = "hello"

	// Act
	str, ok := i.(string)

	// Assert
	assert.True(t, ok)
	assert.Equal(t, "hello", str)
}

func TestChainedMethodCalls(t *testing.T) {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}

	greeting := person.Greet()           // want `// Act statement expected`
	assert.Contains(t, greeting, "Josh") // want `// Assert statement expected`
}

func TestChainedMethodCallsValid(t *testing.T) {
	// Arrange
	person := &Person{Name: "Josh", Age: 30}

	// Act
	greeting := person.Greet()

	// Assert
	assert.Contains(t, greeting, "Josh")
}
