package validmultilinecomments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello_LongBelowArrange(t *testing.T) {
	// Arrange
	// All the things
	name := "Josh"

	// Act
	actual := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_LongAboveArrange(t *testing.T) {
	// All the things
	// Arrange
	name := "Josh"

	// Act
	actual := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_LongBelowAct(t *testing.T) {
	// Arrange
	name := "Josh"

	// Act
	// The part
	actual := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_LongAboveAct(t *testing.T) {
	// Arrange
	name := "Josh"

	// The part
	// Act
	actual := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_LongBelowAssert(t *testing.T) {
	// Arrange
	name := "Josh"

	// Act
	actual := SayHello(name)

	// Assert
	// Everything
	assert.Equal(t, "Hello Josh", actual)
}

func TestSayHello_LongAboveAssert(t *testing.T) {
	// Arrange
	name := "Josh"

	// The part
	// Act
	actual := SayHello(name)

	// Everything
	// Assert
	assert.Equal(t, "Hello Josh", actual)
}
