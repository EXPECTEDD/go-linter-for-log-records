package plugin

import (
	"linter/checks/checkislower"
	"linter/checks/checklanguage"
	"linter/checks/checksensitivedata"
	"linter/checks/checkspecialcharactersoremoji"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

type plugin struct {
	analyzers []*analysis.Analyzer
}

func init() {
	register.Plugin("checkislower", newPlugin(checkislower.Analyzer))
	register.Plugin("checklanguage", newPlugin(checklanguage.Analyzer))
	register.Plugin("checksensitivedata", newPlugin(checksensitivedata.Analyzer))
	register.Plugin("checkspecialcharactersoremoji", newPlugin(checkspecialcharactersoremoji.Analyzer))
}

func newPlugin(a *analysis.Analyzer) func(any) (register.LinterPlugin, error) {
	return func(settings any) (register.LinterPlugin, error) {
		return &plugin{
			analyzers: []*analysis.Analyzer{a},
		}, nil
	}
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return p.analyzers, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
