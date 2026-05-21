package triplea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMisspelledArrange(t *testing.T) {
	// Arrang
	name := "Josh" // want `// Arrange statement expected`

	// Act
	result := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", result)
}

func TestMisspelledAct(t *testing.T) {
	// Axt
	result := SayHello("Josh") // want `// Act statement expected`

	// Assert
	assert.Equal(t, "Hello Josh", result)
}

func TestMisspelledAssert(t *testing.T) {
	// Act
	result := SayHello("Josh")

	// Asert
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestLowercaseAct(t *testing.T) {
	// act
	result := SayHello("Josh") // want `// Act statement expected`

	// assert
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestUppercaseAct(t *testing.T) {
	// ACT
	result := SayHello("Josh") // want `// Act statement expected`

	// ASSERT
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestMixedCaseAct(t *testing.T) {
	// aRrAnGe
	name := "Josh" // want `// Arrange statement expected`

	// AcT
	result := SayHello(name) // want `// Act statement expected`

	// aSsErT
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestActWithExtraText(t *testing.T) {
	// Act - calling the function
	result := SayHello("Josh") // want `// Act statement expected`

	// Assert - verify result
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestActWithColon(t *testing.T) {
	// Act:
	result := SayHello("Josh") // want `// Act statement expected`

	// Assert:
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestActCommentTooFarAbove(t *testing.T) {
	// Act

	// Some other comment
	result := SayHello("Josh") // want `// Act statement expected`

	// Assert
	assert.Equal(t, "Hello Josh", result)
}

func TestActCommentMultipleLinesAbove(t *testing.T) {
	// This is a long comment
	// that spans multiple lines
	// Act

	result := SayHello("Josh") // want `// Act statement expected`

	// Assert
	assert.Equal(t, "Hello Josh", result)
}
