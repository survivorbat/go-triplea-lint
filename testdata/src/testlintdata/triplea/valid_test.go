package triplea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	t.Parallel()
	// Arrange
	name := "Josh"

	// Act
	actual := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayGoodBye(t *testing.T) {
	// Arrange
	name := "Josh"

	// Act
	actual := SayGoodbye(name)

	// Assert
	assert.Equal(t, "Goodbye Josh", actual)
}

func TestSayGoodMorning(t *testing.T) {
	// Act
	actual := SayGoodMorning("Josh")

	// Assert
	assert.Equal(t, "Good morning Josh", actual)
}

func TestSayGoodLateMorning(t *testing.T) {
	t.Parallel()
	// Act
	actual := SayGoodLateMorning("Josh")

	// Assert
	assert.Equal(t, "Good late morning Josh", actual)
}

func TestSayGoodEvening(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good evening Josh"},
		"Anne": {input: "Anne", expected: "Good evening Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			// Act
			actual := SayGoodEvening(testData.input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

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

func TestSayGoodDay(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good day Josh"},
		"Anne": {input: "Anne", expected: "Good day Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			input := testData.input

			// Act
			actual := SayGoodDay(input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestSayGoodNight(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good night Josh"},
		"Anne": {input: "Anne", expected: "Good night Anne"},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Arrange
			input := testData.input

			// Act
			actual := SayGoodNight(input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}
