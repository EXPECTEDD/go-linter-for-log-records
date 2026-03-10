package checklanguage

import (
	"go/ast"
	"linter/checks"
	"slices"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "checkLanguage",
	Doc:  "report if the language is not eng",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			selector, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			pack, ok := selector.X.(*ast.Ident)
			if !ok {
				return true
			}

			if pack.Name == checks.ExpPackageName && slices.Contains(checks.ExpLevelsNames, selector.Sel.Name) {
				text := call.Args[0].(*ast.BasicLit)
				if len(text.Value) > 2 {
					for _, r := range text.Value[1 : len(text.Value)-1] {
						if unicode.IsLetter(r) && !checkIsEng(r) {
							pass.Reportf(call.Pos(), "the log must contain only English characters - %s", text.Value)
							return true
						}
					}
				}
			}

			return true
		})
	}
	return nil, nil
}

func checkIsEng(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}
	return false
}
