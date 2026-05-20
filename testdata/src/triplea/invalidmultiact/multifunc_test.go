package triplea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple functions

func TestSayHello_MultipleAct(t *testing.T) {
	// Arrange
	name1 := "Josh"
	name2 := "Anne"

	// Act
	actual1 := SayHello(name1)
	actual2 := SayHello(name2)

	// Assert
	assert.Equal(t, "Hello Josh", actual1)
	assert.Equal(t, "Hello Anne", actual2)
}

func TestSayHello_MultipleActInParallel(t *testing.T) {
	t.Parallel()
	// Arrange
	name1 := "Josh"
	name2 := "Anne"

	// Act
	actual1 := SayHello(name1)
	actual2 := SayHello(name2)

	// Assert
	assert.Equal(t, "Hello Josh", actual1)
	assert.Equal(t, "Hello Anne", actual2)
}

// Simple functions, table-driven

func TestSayHello_MultipleActInTable(t *testing.T) {
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
			name1 := testData.name1
			name2 := testData.name2

			// Act
			result1 := SayHello(name1)
			result2 := SayHello(name2)

			// Assert
			assert.Equal(t, testData.expected1, result1)
			assert.Equal(t, testData.expected2, result2)
		})
	}
}

func TestSayHello_MultipleActInTableInParallel(t *testing.T) {
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
			name1 := testData.name1
			name2 := testData.name2

			// Act
			result1 := SayHello(name1)
			result2 := SayHello(name2)

			// Assert
			assert.Equal(t, testData.expected1, result1)
			assert.Equal(t, testData.expected2, result2)
		})
	}
}
