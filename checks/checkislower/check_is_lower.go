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
			call, ok := n.(*ast.CallExpr) // проверяем вызов функции
			if !ok {
				return true
			}

			selector, ok := call.Fun.(*ast.SelectorExpr) // проверяем что вызов имеет две части (имеет .)
			if !ok {
				return true
			}

			// Получаем информацию о типе в дереве
			obj := pass.TypesInfo.Uses[selector.Sel]
			if obj == nil {
				return true
			}

			// Проверяем что это функция и получаем ее
			fn, ok := obj.(*types.Func)
			if !ok {
				return true
			}

			// Получаем сигнатуру функции
			sig := fn.Type().(*types.Signature)
			if sig == nil {
				return true
			}

			// Получаем ресеивер функции, пример func (l *Logger) <- receiver
			recv := sig.Recv()
			if recv == nil {
				return true
			}

			// Получаем тип
			loggerType := recv.Type()
			if loggerType == nil {
				return true
			}

			typeStrs := strings.Split(loggerType.String(), "/")

			if slices.Contains(checks.ExpPackageName, typeStrs[len(typeStrs)-1]) && slices.Contains(checks.ExpLevelsNames, selector.Sel.Name) {
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
