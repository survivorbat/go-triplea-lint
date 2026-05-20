package validfunctiontables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// No parallel

func TestSayHello(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Hello Josh"},
		"Anne": {input: "Anne", expected: "Hello Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			input := testData.input

			// Act
			actual := SayHello(input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestSayHello_ReturnsExpectedMessage(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Hello Josh"},
		"Anne": {input: "Anne", expected: "Hello Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			input := testData.input

			// Act
			actual := SayHello(input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

// Parallel

func TestSayGoodbye(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Goodbye Josh"},
		"Anne": {input: "Anne", expected: "Goodbye Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Arrange
			input := testData.input

			// Act
			actual := SayGoodbye(input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestSayGoodbye_ReturnsExpectedMessage(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Goodbye Josh"},
		"Anne": {input: "Anne", expected: "Goodbye Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Arrange
			input := testData.input

			// Act
			actual := SayGoodbye(input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

// No arrange, no parallel

func TestSayGoodMorning(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good morning Josh"},
		"Anne": {input: "Anne", expected: "Good morning Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			// Act
			actual := SayGoodMorning(testData.input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestSayGoodMorning_ReturnsExpectedMessage(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good morning Josh"},
		"Anne": {input: "Anne", expected: "Good morning Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			// Act
			actual := SayGoodMorning(testData.input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

// No arrange, parallel

func TestSayGoodAfternoon(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good afternoon Josh"},
		"Anne": {input: "Anne", expected: "Good afternoon Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Act
			actual := SayGoodAfternoon(testData.input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestSayGoodAfternoon_ReturnsExpectedMessage(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good afternoon Josh"},
		"Anne": {input: "Anne", expected: "Good afternoon Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Act
			actual := SayGoodAfternoon(testData.input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}
