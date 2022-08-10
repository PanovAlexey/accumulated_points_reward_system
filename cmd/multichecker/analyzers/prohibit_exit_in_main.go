package analyzers

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"strings"
)

var AnalyzerProhibitExitInMain = &analysis.Analyzer{
	Name: "prohibitexit",
	Doc: "Prohibit method os.Exit calling in main. " +
		"The analyzer takes as input the path to the directory with GO files.",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		if strings.ToLower(f.Name.Name) == "main" {
			ast.Inspect(f, func(node ast.Node) bool {
				if funcMain, ok := node.(*ast.FuncDecl); ok && funcMain.Name.Name == "main" {
					check(node, pass)
				}

				return true
			})
		}
	}

	return nil, nil
}

func check(node ast.Node, pass *analysis.Pass) {
	ast.Inspect(node, func(node ast.Node) bool {
		if node != nil {
			if call, ok := node.(*ast.CallExpr); ok {
				if selector, ok := call.Fun.(*ast.SelectorExpr); ok {
					if pkg, ok := selector.X.(*ast.Ident); ok && pkg.String() == "os" && selector.Sel.Name == "Exit" {
						pass.Reportf(node.Pos(), "os.Exit call in main.go")

						return true
					}
				}
			}
		}

		return true
	})
}
