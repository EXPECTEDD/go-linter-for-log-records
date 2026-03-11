package plugin

import (
	"linter/checks/checkislower"
	"linter/checks/checklanguage"
	"linter/checks/checksensitivedata"
	"linter/checks/checkspecialcharactersoremoji"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

type plugin struct{}

func init() {
	register.Plugin("mylinter", New)
}

func New(settings any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

var _ register.LinterPlugin = (*plugin)(nil)

func (*plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		checkislower.Analyzer,
		checklanguage.Analyzer,
		checksensitivedata.Analyzer,
		checkspecialcharactersoremoji.Analyzer,
	}, nil
}

func (*plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
