package checkspecialcharactersoremoji

import (
	"go/ast"
	"linter/checks"
	"slices"
	"unicode"

	"github.com/forPelevin/gomoji"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "checkSpecialCharactersOrEmoji",
	Doc:  "report if there are special characters or emojis",
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
				text := call.Args[0].(*ast.BasicLit)
				if gomoji.ContainsEmoji(text.Value[1 : len(text.Value)-1]) {
					pass.Reportf(call.Pos(), "the log contains emoji - %s", text.Value)
					return true
				}
				textRune := []rune(text.Value[1 : len(text.Value)-1])
				for _, r := range textRune {
					if checkSpecialCharactersOrEmoji(r) {
						pass.Reportf(call.Pos(), "the log contains special character - %s", text.Value)
						return true
					}
				}
			}

			return true
		})
	}
	return nil, nil
}

func checkSpecialCharactersOrEmoji(r rune) bool {
	if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
		return true
	}
	return false
}
