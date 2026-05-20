package validmethods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Shortened name

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
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayHello()

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
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayHello()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

// No parallel

func TestMyService_SayHello(t *testing.T) {
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
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayHello()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestMyService_SayHello_ReturnsExpectedMessage(t *testing.T) {
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
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayHello()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

// Parallel

func TestMyService_SayGoodbye(t *testing.T) {
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
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayGoodbye()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestMyService_SayGoodbye_ReturnsExpectedMessage(t *testing.T) {
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
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayGoodbye()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

// No arrange, no parallel

func TestMyService_SayGoodMorning(t *testing.T) {
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
			actual := (&MyService{Name: testData.input}).SayGoodMorning()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestMyService_SayGoodMorning_ReturnsExpectedMessage(t *testing.T) {
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
			actual := (&MyService{Name: testData.input}).SayGoodMorning()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

// No arrange, parallel

func TestMyService_SayGoodAfternoon(t *testing.T) {
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
			actual := (&MyService{Name: testData.input}).SayGoodAfternoon()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestMyService_SayGoodAfternoon_ReturnsExpectedMessage(t *testing.T) {
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
			actual := (&MyService{Name: testData.input}).SayGoodAfternoon()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}
