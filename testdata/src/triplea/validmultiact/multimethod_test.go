package triplea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple methods

func TestPerson_SayHello_MultipleAct(t *testing.T) {
	// Arrange
	person1 := &Person{Name: "Josh"}
	person2 := &Person{Name: "Anne"}

	// Act
	actual1 := person1.SayHello()
	actual2 := person2.SayHello()

	// Assert
	assert.Equal(t, "Hello Josh", actual1)
	assert.Equal(t, "Hello Anne", actual2)
}

func TestPerson_SayHello_MultipleActInParallel(t *testing.T) {
	t.Parallel()
	// Arrange
	person1 := &Person{Name: "Josh"}
	person2 := &Person{Name: "Anne"}

	// Act
	actual1 := person1.SayHello()
	actual2 := person2.SayHello()

	// Assert
	assert.Equal(t, "Hello Josh", actual1)
	assert.Equal(t, "Hello Anne", actual2)
}

// Simple methods, table-driven

func TestPerson_SayHello_MultipleActInTable(t *testing.T) {
	tests := map[string]struct {
		name1 string
		name2 string

		expected1 string
		expected2 string
	}{
		"Josh and Anne":     {name1: "Josh", name2: "Anne", expected1: "Hello Josh", expected2: "Hello Anne"},
		"Jeniffer and Mila": {name1: "Jeniffer", name2: "Mila", expected1: "Hello Jeniffer", expected2: "Hello Mila"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			person1 := &Person{Name: testData.name1}
			person2 := &Person{Name: testData.name2}

			// Act
			result1 := person1.SayHello()
			result2 := person2.SayHello()

			// Assert
			assert.Equal(t, testData.expected1, result1)
			assert.Equal(t, testData.expected2, result2)
		})
	}
}

func TestPerson_SayHello_MultipleActInTableInParallel(t *testing.T) {
	tests := map[string]struct {
		name1 string
		name2 string

		expected1 string
		expected2 string
	}{
		"Josh and Anne":     {name1: "Josh", name2: "Anne", expected1: "Hello Josh", expected2: "Hello Anne"},
		"Jeniffer and Mila": {name1: "Jeniffer", name2: "Mila", expected1: "Hello Jeniffer", expected2: "Hello Mila"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Arrange
			person1 := &Person{Name: testData.name1}
			person2 := &Person{Name: testData.name2}

			// Act
			result1 := person1.SayHello()
			result2 := person2.SayHello()

			// Assert
			assert.Equal(t, testData.expected1, result1)
			assert.Equal(t, testData.expected2, result2)
		})
	}
}
