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
			call, ok := n.(*ast.CallExpr) // проверяем вызов функции
			if !ok {
				return true
			}

			selector, ok := call.Fun.(*ast.SelectorExpr) // проверяем что вызов имеет две части (имеет .)
			if !ok {
				return true
			}

			pack, ok := selector.X.(*ast.Ident) // получаем название пакета
			if !ok {
				return true
			}

			if pack.Name == checks.ExpPackageName && slices.Contains(checks.ExpLevelsNames, selector.Sel.Name) {
				resBasLit := make([]*ast.BasicLit, 0)

				for _, arg := range call.Args {
					walkExpr(arg, &resBasLit)
				}
				for _, b := range resBasLit {
					if len(b.Value) > 2 {
						for _, r := range b.Value[1 : len(b.Value)-1] {
							if unicode.IsLetter(r) && !checkIsEng(r) {
								pass.Reportf(call.Pos(), "the log must contain only English characters - %s", b.Value)
								return true
							}
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

func walkExpr(expr ast.Expr, result *[]*ast.BasicLit) {
	switch v := expr.(type) {
	case *ast.BasicLit:
		*result = append(*result, v)
	case *ast.BinaryExpr:
		walkExpr(v.X, result)
		walkExpr(v.Y, result)
	case *ast.CallExpr:
		walkExpr(v.Fun, result)
		for _, arg := range v.Args {
			walkExpr(arg, result)
		}
	}
}
