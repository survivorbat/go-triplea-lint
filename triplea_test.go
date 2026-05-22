package linters

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := map[string]struct {
		patterns string
		options  map[string]string
	}{
		"triplea": {
			patterns: "triplea",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			// Arrange
			a := Analyzer()

			for k, v := range test.options {
				err := a.Flags.Set(k, v)
				if err != nil {
					t.Fatal(err)
				}
			}

			// Act & Assert
			analysistest.Run(t, analysistest.TestData(), a, test.patterns)
		})
	}
}
