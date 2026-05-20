package false

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMain should not be linted
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	// Setup code
}

func teardown() {
	// Teardown code
}

// Helper functions should not be linted

func helperFunction(t *testing.T) string {
	t.Helper()
	result := SayHello("Josh")
	return result
}

func anotherHelper(name string) string {
	greeting := SayHello(name)
	return greeting
}

func createTestPerson() *Person {
	person := &Person{Name: "Josh", Age: 30}
	return person
}

// Private functions (lowercase) should not be linted

func testHelper(t *testing.T) {
	t.Helper()
	name := "Josh"
	result := SayHello(name)
	assert.Equal(t, "Hello Josh", result)
}

func validateResult(t *testing.T, result string) {
	assert.NotEmpty(t, result)
	assert.Contains(t, result, "Hello")
}

// Benchmark functions should not be linted

func BenchmarkSayHello(b *testing.B) {
	for b.Loop() {
		SayHello("Josh")
	}
}

func BenchmarkNewPerson(b *testing.B) {
	for b.Loop() {
		NewPerson("Josh", 30)
	}
}

// Example functions should not be linted

func ExampleSayHello() {
	result := SayHello("Josh")
	println(result)
}

func ExamplePerson_Greet() {
	person := &Person{Name: "Josh", Age: 30}
	result := person.Greet()
	println(result)
}

// Fuzz functions should not be linted

func FuzzSayHello(f *testing.F) {
	f.Add("Josh")
	f.Fuzz(func(t *testing.T, name string) {
		result := SayHello(name)
		assert.Contains(t, result, name)
	})
}

// Methods that happen to start with Test but aren't test functions

type TestHelper struct {
	value string
}

func (h *TestHelper) TestMethod() string {
	result := SayHello(h.value)
	return result
}

func (h *TestHelper) TestAnotherMethod() {
	h.value = "updated"
}

// Table setup functions

func getTestCases() map[string]struct {
	input    string
	expected string
} {
	return map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Hello Josh"},
		"Anne": {input: "Anne", expected: "Hello Anne"},
	}
}

func generateTestData() []Person {
	return []Person{
		{Name: "Josh", Age: 30},
		{Name: "Anne", Age: 25},
	}
}

// Assertion helpers

func assertPersonEqual(t *testing.T, expected, actual *Person) {
	t.Helper()
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Age, actual.Age)
}

func assertStringContains(t *testing.T, s, substr string) {
	t.Helper()
	assert.Contains(t, s, substr)
}

// Cleanup functions

func cleanupTest(t *testing.T) {
	t.Cleanup(func() {
		// cleanup logic
	})
}

func resetState() {
	// reset some state
}
