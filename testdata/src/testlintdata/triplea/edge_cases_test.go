package triplea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	// Empty test body
}

func TestOnlyComments(t *testing.T) {
	// This test has only comments
	// No actual statements
}

func TestNoAssertionsAfterAct(t *testing.T) {
	// Act
	_ = SayHello("Josh")
	// Test ends without assertions
}

func TestNoMatchingFunctionCall(t *testing.T) {
	// Act
	result := "some hardcoded value" // want `// Act statement expected`

	// Assert
	assert.Equal(t, "some hardcoded value", result)
}

func TestMultipleActCandidates(t *testing.T) {
	// Arrange
	name := "Josh"

	// Act
	result1 := SayHello(name)
	result2 := SayHello("Anne")

	// Assert
	assert.Equal(t, "Hello Josh", result1)
	assert.Equal(t, "Hello Anne", result2)
}

func TestEarlyReturn(t *testing.T) {
	// Arrange
	name := ""

	if name == "" {
		return
	}

	// Act
	result := SayHello(name)

	// Assert
	assert.Equal(t, "Hello ", result)
}

func TestMultipleAssertions(t *testing.T) {
	// Act
	result := SayHello("Josh")

	// Assert
	assert.NotEmpty(t, result)
	assert.Contains(t, result, "Josh")
	assert.Equal(t, "Hello Josh", result)
}

func TestWithDefer(t *testing.T) {
	// Arrange
	defer func() {
		// cleanup
	}()

	// Act
	result := SayHello("Josh")

	// Assert
	assert.Equal(t, "Hello Josh", result)
}

func TestNestedRun(t *testing.T) {
	t.Run("outer", func(t *testing.T) {
		t.Run("inner", func(t *testing.T) {
			// Act
			result := SayHello("Josh")

			// Assert
			assert.Equal(t, "Hello Josh", result)
		})
	})
}

func TestNestedRunMissingComments(t *testing.T) {
	t.Run("outer", func(t *testing.T) {
		t.Run("inner", func(t *testing.T) {
			result := SayHello("Josh") // want `// Act statement expected`
			assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
		})
	})
}

func TestNestedRunOuterMissingComments(t *testing.T) {
	t.Run("outer", func(t *testing.T) {
		result := SayHello("Josh") // want `// Act statement expected`
		assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
	})
}

func TestNestedRunInnerOnlyMissingComments(t *testing.T) {
	t.Run("outer", func(t *testing.T) {
		// Act
		result1 := SayHello("Josh")

		// Assert
		assert.Equal(t, "Hello Josh", result1)

		t.Run("inner", func(t *testing.T) {
			result2 := SayGoodbye("Anne") // want `// Act statement expected`
			assert.Equal(t, "Goodbye Anne", result2) // want `// Assert statement expected`
		})
	})
}

func TestTripleNestedRun(t *testing.T) {
	t.Run("level1", func(t *testing.T) {
		t.Run("level2", func(t *testing.T) {
			t.Run("level3", func(t *testing.T) {
				// Act
				result := SayHello("Josh")

				// Assert
				assert.Equal(t, "Hello Josh", result)
			})
		})
	})
}

func TestTripleNestedRunMissingComments(t *testing.T) {
	t.Run("level1", func(t *testing.T) {
		t.Run("level2", func(t *testing.T) {
			t.Run("level3", func(t *testing.T) {
				result := SayHello("Josh") // want `// Act statement expected`
				assert.Equal(t, "Hello Josh", result) // want `// Assert statement expected`
			})
		})
	})
}

func TestRunWithNilFunction(t *testing.T) {
	// Linter should not panic on nil function
	t.Run("nil subtest", nil)
}

func TestRunWithNilFunctionInValid(t *testing.T) {
	// Act
	result := SayHello("Josh")

	// Assert
	assert.Equal(t, "Hello Josh", result)

	// Linter should not panic on nil function
	t.Run("nil subtest", nil)
}

func TestRunWithVariableFunction(t *testing.T) {
	var fn func(*testing.T)
	// Linter should not panic on variable function
	t.Run("variable subtest", fn)
}
