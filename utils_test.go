package linters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasLineWithPrefix(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		prefix   string
		expected bool
	}{
		// No matches
		"Empty string": {
			prefix:   "Arrange",
			expected: false,
		},
		"no match on single line": {
			input:    "This does not contain arrange",
			prefix:   "Act",
			expected: false,
		},
		"no match on multiple lines": {
			input: `These are
multiple lines that
don't have the prefix`,
			prefix:   "Assert",
			expected: false,
		},

		// Matches
		"Arrange on first line": {
			input: `Arrange on the first line
and the other lines
aren't that interesting`,
			prefix:   "Arrange",
			expected: true,
		},
		"Act on middle line": {
			input: `So we're actually going to
Act now and
run the test`,
			prefix:   "Act",
			expected: true,
		},
		"Assert on last line": {
			input: `Now we're going to
do the
Assert phase`,
			prefix:   "Assert",
			expected: true,
		},

		// Comment variants
		"// Act": {
			input:    "Act",
			prefix:   "Act",
			expected: true,
		},
		"/// Act": {
			input:    "/ Act",
			prefix:   "Act",
			expected: true,
		},
		"//// Act": {
			input:    "// Act",
			prefix:   "Act",
			expected: true,
		},
		"//* Act": {
			input:    "* Act",
			prefix:   "Act",
			expected: true,
		},
		"//** Act": {
			input:    "** Act",
			prefix:   "Act",
			expected: true,
		},
		"///***/ Act": {
			input:    "/***/ Act",
			prefix:   "Act",
			expected: true,
		},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Act
			actual := hasLineWithPrefix(testData.input, testData.prefix)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestActCandidates(t *testing.T) {
	t.Parallel()

	tests := map[string][][]string{
		"TestSay": {
			{"say"},
		},
		"TestSayMyName": {
			{"saymyname"},
			{"saymyname", "saymy", "say", "myname", "my", "name"},
		},
		"TestService_Say": {
			{"service_say"},
			{"say", "service"},
		},
		"TestService_SayMyName": {
			{"service_saymyname"},
			{"saymyname", "service"},
			{"saymyname", "saymy", "say", "myname", "my", "name"},
		},
		"TestMyService_Name": {
			{"myservice_name"},
			{"name", "myservice"},
			{"myservice", "my", "service"},
		},
		"TestMyService_SayMyName_DoesAnotherThing": {
			{"myservice_saymyname_doesanotherthing"},
			{"saymyname", "myservice"},
			{"saymyname", "saymy", "say", "myname", "my", "name"},
			{"myservice", "my", "service"},
		},
	}

	for name, expected := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Act
			actual := actCandidates(name)

			// Assert
			assert.Equal(t, expected, actual)
		})
	}
}

func TestMatchWords(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    []string
		expected []string
	}{
		"": {
			input:    []string{},
			expected: []string{},
		},
		"a": {
			input:    []string{"a"},
			expected: []string{"a"},
		},
		"IAmVeryHappy": {
			input: []string{"I", "Am", "Very", "Happy"},
			expected: []string{
				"IAmVeryHappy",

				"IAmVery",
				"IAm",
				"I",

				"AmVeryHappy",
				"AmVery",
				"Am",

				"VeryHappy",
				"Very",

				"Happy",
			},
		},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Act
			actual := matchWords(testData.input)

			// Assert
			assert.Equal(t, testData.expected, actual)
		})
	}
}

func TestSplitWords(t *testing.T) {
	t.Parallel()

	tests := map[string][]string{
		"":  {},
		"a": {"a"},
		"B": {"B"},

		"helloworld": {"helloworld"},
		"helloWorld": {"hello", "World"},
		"HelloWorld": {"Hello", "World"},
		"hElLoWoRlD": {"h", "El", "Lo", "Wo", "Rl", "D"},
		"HeLlOwOrLd": {"He", "Ll", "Ow", "Or", "Ld"},

		"HELLOWORLD": {"HELLOWORLD"},
		"helloWORLD": {"hello", "WORLD"},
		"myHTTPTest": {"my", "HTTP", "Test"},
		"HTTPTest":   {"HTTP", "Test"},
		"testHTTP":   {"test", "HTTP"},
	}

	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			// Act
			actual := splitWords(input)

			// Assert
			assert.Equal(t, expected, actual)
		})
	}
}
