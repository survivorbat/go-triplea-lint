package linters

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

const rootTestDir = "testdata/src/triplea"

func TestPlugin_BuildAnalyzers_AnalysesCorrectly(t *testing.T) {
	t.Parallel()

	for _, dirName := range getTestDirectories(t) {
		t.Run(dirName, func(t *testing.T) {
			// Arrange
			plugin, err := New(nil)
			require.NoError(t, err)

			testDir := path.Join("triplea", dirName, "...")

			// Act
			actual, err := plugin.BuildAnalyzers()

			// Assert
			require.NoError(t, err)
			require.Len(t, actual, 1)

			analysistest.Run(t, analysistest.TestData(), actual[0], testDir)
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
