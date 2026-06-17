// This must be package main
package main

import (
	linters "github.com/survivorbat/go-triplea-lint"
	"golang.org/x/tools/go/analysis"
)

func New(any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{linters.Analyzer()}, nil
}
