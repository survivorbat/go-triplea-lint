package linters

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	// Arrange
	analyzer := Analyzer()
	directory := filepath.Join("testdata", "src", "testlintdata", "triplea")

	// Act & Assert
	analysistest.Run(t, directory, analyzer, "missing_comments_test.go", "funcs.go")
}
