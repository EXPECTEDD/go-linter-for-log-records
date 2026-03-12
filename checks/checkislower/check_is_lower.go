package checkislower

import (
	"go/ast"
	"go/token"
	"go/types"
	"linter/checks"
	"slices"
	"strings"
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
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			selector, ok := call.Fun.(*ast.SelectorExpr)
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

			pkg := obj.Pkg()
			if pkg == nil {
				return true
			}

			isFunc := false
			if (pkg.Path() == "log/slog" || pkg.Path() == "go.uber.org/zap") && slices.Contains(checks.ExpLevelsNames, selector.Sel.Name) {
				isFunc = true
			}

			isLoggerMethod := false
			recv := sig.Recv()
			if recv != nil {
				loggerType := recv.Type()
				typeStrs := strings.Split(loggerType.String(), "/")
				if slices.Contains(checks.ExpPackageName, typeStrs[len(typeStrs)-1]) && slices.Contains(checks.ExpLevelsNames, selector.Sel.Name) {
					isLoggerMethod = true
				}
			}

			// Если это вызов slog.Debug() или logger.Debug()
			if isFunc || isLoggerMethod {
				if len(call.Args) == 0 {
					return true
				}

				textRune := []rune(walkExpr(call.Args[0]))
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
		if v.Kind == token.STRING {
			return v.Value
		}
	case *ast.BinaryExpr:
		return walkExpr(v.X)
	case *ast.ParenExpr:
		return walkExpr(v.X)
	}
	return ""
}
