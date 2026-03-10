package checksensitivedata

import (
	"go/ast"
	"go/types"
	"linter/checks"
	"slices"
	"strings"

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

			obj := pass.TypesInfo.Uses[selector.Sel]
			if obj == nil {
				return true
			}

			fn, ok := obj.(*types.Func)
			if !ok {
				return true
			}

			sig := fn.Type().(*types.Signature)
			if sig == nil {
				return true
			}

			recv := sig.Recv()
			if recv == nil {
				return true
			}

			loggerType := recv.Type()
			if loggerType == nil {
				return true
			}

			typeStrs := strings.Split(loggerType.String(), "/")

			if slices.Contains(checks.ExpPackageName, typeStrs[len(typeStrs)-1]) && slices.Contains(checks.ExpLevelsNames, selector.Sel.Name) {
				result := make([]*ast.Ident, 0)
				for _, arg := range call.Args {
					walkExpr(arg, &result)
				}
				for _, ident := range result {
					if slices.Contains(checks.ListSensetiveData, strings.ToLower(ident.Name)) {
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
	case *ast.ParenExpr:
		walkExpr(v.X, result)
	}
}
