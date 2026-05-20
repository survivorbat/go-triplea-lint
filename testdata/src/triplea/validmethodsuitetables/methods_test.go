package validmethods

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ValidSuite struct {
	suite.Suite
}

// No T

func (v *ValidSuite) TestSayHello() {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Hello Josh"},
		"Anne": {input: "Anne", expected: "Hello Anne"},
	}

	for name, testData := range tests {
		v.Run(name, func() {
			// Arrange
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayHello()

			// Assert
			assert.Equal(v.T(), testData.expected, actual)
		})
	}
}

func (v *ValidSuite) TestSayHello_ReturnsExpectedMessage() {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Hello Josh"},
		"Anne": {input: "Anne", expected: "Hello Anne"},
	}

	for name, testData := range tests {
		v.Run(name, func() {
			// Arrange
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayHello()

			// Assert
			assert.Equal(v.T(), testData.expected, actual)
		})
	}
}

// T

func (v *ValidSuite) TestSayGoodBye() {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Goodbye Josh"},
		"Anne": {input: "Anne", expected: "Goodbye Anne"},
	}

	for name, testData := range tests {
		v.Run(name, func() {
			t := v.T()
			// Arrange
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayGoodbye()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func (v *ValidSuite) TestSayGoodBye_ReturnsExpectedMessage() {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Goodbye Josh"},
		"Anne": {input: "Anne", expected: "Goodbye Anne"},
	}

	for name, testData := range tests {
		v.Run(name, func() {
			t := v.T()
			// Arrange
			svc := &MyService{Name: testData.input}

			// Act
			actual := svc.SayGoodbye()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

// No arrange, no T

func (v *ValidSuite) TestSayGoodMorning() {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good morning Josh"},
		"Anne": {input: "Anne", expected: "Good morning Anne"},
	}

	for name, testData := range tests {
		v.Run(name, func() {
			// Act
			actual := (&MyService{Name: testData.input}).SayGoodMorning()

			// Assert
			assert.Equal(v.T(), testData.expected, actual)
		})
	}
}

func (v *ValidSuite) TestSayGoodMorning_ReturnsExpectedMessage() {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good morning Josh"},
		"Anne": {input: "Anne", expected: "Good morning Anne"},
	}

	for name, testData := range tests {
		v.Run(name, func() {
			// Act
			actual := (&MyService{Name: testData.input}).SayGoodMorning()

			// Assert
			assert.Equal(v.T(), testData.expected, actual)
		})
	}
}

// No arrange, T

func (v *ValidSuite) TestSayGoodAfternoon() {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good afternoon Josh"},
		"Anne": {input: "Anne", expected: "Good afternoon Anne"},
	}

	for name, testData := range tests {
		v.Run(name, func() {
			t := v.T()
			// Act
			actual := (&MyService{Name: testData.input}).SayGoodAfternoon()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func (v *ValidSuite) TestSayGoodAfternoon_ReturnsExpectedMessage() {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"Josh": {input: "Josh", expected: "Good afternoon Josh"},
		"Anne": {input: "Anne", expected: "Good afternoon Anne"},
	}

	for name, testData := range tests {
		v.Run(name, func() {
			t := v.T()
			// Act
			actual := (&MyService{Name: testData.input}).SayGoodAfternoon()

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestValidSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ValidSuite))
}
