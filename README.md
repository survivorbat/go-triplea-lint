# 🎗️ Triple-A Lint

This linter aims to enforce the use of `Arrange/Act/Assert` or `Given/When/Then`
comments in tests.

## 🤔 Why?

Tests usually follow a simple pattern: Setup, execute, verify.
For simple tests this pattern is obvious, in more complicated tests it
becomes harder to follow _what_ is being tested.

One option to solve this problem is to add comments to tests, such as:

```go
package printer

import (
  "testing"

  "github.com/survivorbat/go-triplea-lint/document"

  "github.com/stretchr/testify/require"
  "github.com/stretchr/testify/assert"
)

func TestPrinter_Print_PrintsDocument(t *testing.T) {
  t.Parallel()
  // Arrange
  printer := New()
  document, err := document.New("...")
  require.NoError(t, err)

  // Act
  jobID, err := printer.Print(t.Context())

  // Assert
  require.NoError(t, err)
  assert.NotEmpty(t, jobID)
}
```

This linter aims to enforce that way of working.

## ⬇️ Installation: Coming soon

## 📋 Usage: Coming soon

## 🔭 Plans

- Add more test cases from [todos.txt](./todos.txt)
- Improve act election
- Cleanup code in `triplea.go`, it's too nested
- Lint this codebase
- Document test command
