package linters

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("triplea", New)
}

type Plugin struct{}

func New(any) (register.LinterPlugin, error) {
	return new(Plugin), nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "triplea",
			Doc:  "Enforce the use of Arrange/Act/Assert comments in tests",
			Run:  run,
		},
	}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
