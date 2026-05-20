package validfunctions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// No parallel

func TestSayHello(t *testing.T) {
	// Arrange
	name := "Josh"

	// Act
	actual := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_ReturnsExpectedMessage(t *testing.T) {
	// Arrange
	name := "Josh"

	// Act
	actual := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

// Parallel

func TestSayGoodBye(t *testing.T) {
	t.Parallel()
	// Arrange
	name := "Josh"

	// Act
	actual := SayGoodbye(name)

	// Assert
	assert.Equal(t, "Goodbye Josh", actual)
}

func TestSayGoodBye_ReturnsExpectedMessage(t *testing.T) {
	t.Parallel()
	// Arrange
	name := "Josh"

	// Act
	actual := SayGoodbye(name)

	// Assert
	assert.Equal(t, "Goodbye Josh", actual)
}

// No arrange, no parallel

func TestSayGoodMorning(t *testing.T) {
	// Act
	actual := SayGoodMorning("Josh")

	// Assert
	assert.Equal(t, "Good morning Josh", actual)
}

func TestSayGoodMorning_ReturnsExpectedMessage(t *testing.T) {
	// Act
	actual := SayGoodMorning("Josh")

	// Assert
	assert.Equal(t, "Good morning Josh", actual)
}

// No arrange, parallel

func TestSayGoodAfternoon(t *testing.T) {
	t.Parallel()
	// Act
	actual := SayGoodAfternoon("Josh")

	// Assert
	assert.Equal(t, "Good afternoon Josh", actual)
}

func TestSayGoodAfternoon_ReturnsExpectedMessage(t *testing.T) {
	t.Parallel()
	// Act
	actual := SayGoodAfternoon("Josh")

	// Assert
	assert.Equal(t, "Good afternoon Josh", actual)
}
