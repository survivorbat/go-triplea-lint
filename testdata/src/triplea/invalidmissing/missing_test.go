package triplea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaveHello(t *testing.T) {
	t.Parallel()
	name := "Josh"                        // want `// Arrange statement expected`
	actual := WaveHello(name)             // want `// Act statement expected`
	assert.Equal(t, "Hello Josh", actual) // want `// Assert statement expected`
}

func TestWaveGoodbye(t *testing.T) {
	name := "Josh" // want `// Arrange statement expected`

	// Act
	actual := WaveGoodbye(name)
	assert.Equal(t, "Goodbye Josh", actual) // want `// Assert statement expected`
}

func TestWaveGoodMorning(t *testing.T) {
	actual := WaveGoodMorning("Josh") // want `// Act statement expected`

	// Assert
	assert.Equal(t, "Good morning Josh", actual)
}

func TestWaveGoodLateMorning(t *testing.T) {
	t.Parallel()
	actual := WaveGoodLateMorning("Josh") // want `// Act statement expected`

	// Assert
	assert.Equal(t, "Good late morning Josh", actual)
}

func TestWaveGoodEvening(t *testing.T) {
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
			actual := WaveGoodEvening(testData.input)  // want `// Act statement expected`
			assert.Equal(t, testData.expected, actual) // want `// Assert statement expected`
		})
	}
}

func TestWaveGoodAfternoon(t *testing.T) {
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
			actual := WaveGoodAfternoon(testData.input) // want `// Act statement expected`
			assert.Equal(t, testData.expected, actual)  // want `// Assert statement expected`
		})
	}
}

func TestWaveGoodDay(t *testing.T) {
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
			input := testData.input                    // want `// Arrange statement expected`
			actual := WaveGoodDay(input)               // want `// Act statement expected`
			assert.Equal(t, testData.expected, actual) // want `// Assert statement expected`
		})
	}
}

func TestWaveGoodNight(t *testing.T) {
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
			input := testData.input                    // want `// Arrange statement expected`
			actual := WaveGoodNight(input)             // want `// Act statement expected`
			assert.Equal(t, testData.expected, actual) // want `// Assert statement expected`
		})
	}
}
