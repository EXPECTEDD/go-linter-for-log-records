package main

import (
	"linter/checks/checkislower"
	"linter/checks/checklanguage"
	"linter/checks/checksensitivedata"
	"linter/checks/checkspecialcharactersoremoji"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		checkislower.Analyzer,
		checklanguage.Analyzer,
		checkspecialcharactersoremoji.Analyzer,
		checksensitivedata.Analyzer,
	)
}
