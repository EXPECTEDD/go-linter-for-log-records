package checksensitivedata

import (
	"go/ast"
	"linter/checks"
	"slices"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "checkSensitiveData",
	Doc:  "report if there is sensitive data",
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
				result := make([]*ast.Ident, 0)
				for _, arg := range call.Args {
					walkExpr(arg, &result)
				}
				for _, ident := range result {
					if slices.Contains(checks.ListSensetiveData, ident.Name) {
						pass.Reportf(call.Pos(), "the log contains sensitive data - %s", ident.Name)
						return true
					}
				}
			}

			return true
		})
	}
	return nil, nil
}

func walkExpr(expr ast.Expr, result *[]*ast.Ident) {
	switch v := expr.(type) {
	case *ast.Ident:
		*result = append(*result, v)
	case *ast.BinaryExpr:
		walkExpr(v.X, result)
		walkExpr(v.Y, result)
	case *ast.CallExpr:
		walkExpr(v.Fun, result)
		for _, arg := range v.Args {
			walkExpr(arg, result)
		}
	case *ast.SelectorExpr:
		*result = append(*result, v.Sel)
		walkExpr(v.X, result)
	}
}
