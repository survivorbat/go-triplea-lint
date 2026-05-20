package linters

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

const rootTestDir = "testdata/src/triplea"

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	for _, dirName := range getTestDirectories(t) {
		t.Run(dirName, func(t *testing.T) {
			// Arrange
			analyzer := Analyzer()

			testDir := path.Join("triplea", dirName, "...")

			// Act & Assert
			analysistest.Run(t, analysistest.TestData(), analyzer, testDir)
		})
	}
}

// getTestDirectories is a wrapper around os.
func getTestDirectories(t *testing.T) []string {
	entries, err := os.ReadDir(rootTestDir)
	require.NoError(t, err)

	result := make([]string, 0, len(entries))

	for _, dirEntry := range entries {
		if dirEntry.IsDir() && dirEntry.Name() != "vendor" {
			result = append(result, dirEntry.Name())
		}
	}

	return result
}
