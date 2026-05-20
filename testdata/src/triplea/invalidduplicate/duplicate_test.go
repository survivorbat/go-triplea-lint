package invalidduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello_DuplicateArrange(t *testing.T) { // want `Duplicate Arrange statement`
	// Arrange

	// Arrange
	name := "Josh"

	// Act
	result := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", result)
}

func TestSayHello_DuplicateAct(t *testing.T) { // want `Duplicate Act statement`
	// Arrange
	name := "Josh"

	// Act

	// Act
	result := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", result)
}

func TestSayHello_DuplicateAssert(t *testing.T) { // want `Duplicate Assert statement`
	// Arrange
	name := "Josh"

	// Act
	result := SayHello(name)

	// Assert

	// Assert
	assert.Equal(t, "Hello Josh", result)
}
