package invalidfunctions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello_MisspelledArrange(t *testing.T) {
	// Arrang
	name := "Josh" // want `// Arrange statement expected`

	// Act
	result := SayHello(name)

	// Assert
	assert.Equal(t, "Hello Josh", result)
}

func TestSayHello_MisspelledAct(t *testing.T) {
	// Axt
	result := SayHello("Josh") // want `// Act statement expected`

	// Assert
	assert.Equal(t, "Hello Josh", result)
}

func TestSayHello_MisspelledAssert(t *testing.T) {
	// Act
	result := SayHello("Josh")

	// Asert
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestSayHello_Lowercase(t *testing.T) {
	// act
	result := SayHello("Josh") // want `// Act statement expected`

	// assert
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestSayHello_Uppercase(t *testing.T) {
	// ACT
	result := SayHello("Josh") // want `// Act statement expected`

	// ASSERT
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestSayHello_MixedCase(t *testing.T) {
	// aRrAnGe
	name := "Josh" // want `// Arrange statement expected`

	// AcT
	result := SayHello(name) // want `// Act statement expected`

	// aSsErT
	assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
}

func TestSayHello_ActCommentTooFarAbove(t *testing.T) {
	// Act

	// Some other comment
	result := SayHello("Josh") // want `// Act statement expected`

	// Assert
	assert.Equal(t, "Hello Josh", result)
}
