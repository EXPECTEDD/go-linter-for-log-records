package checkislower

import (
	"go/ast"
	"linter/checks"
	"slices"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "checkIsLower",
	Doc:  "report if log doesn't start with a lowercase letter",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
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
				textRune := []rune(walkExpr(call.Args[0])) // получаем первый аргумент функции
				if len(textRune) > 2 && unicode.IsUpper(textRune[1]) {
					pass.Reportf(call.Pos(), "log must start with a lowercase letter - %s", string(textRune))
				}
			}

			return true
		})
	}
	return nil, nil
}

func walkExpr(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.BasicLit:
		return v.Value
	case *ast.BinaryExpr:
		return walkExpr(v.X)
	}
	return ""
}
