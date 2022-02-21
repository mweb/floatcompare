// Copyright: (c) 2022, Mathias Weber (mweb@gmx.ch)

package floatcompare

import (
	"bytes"
	"flag"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type floatcompare struct {
	equalOnly bool
	skipTests bool
}

func NewAnalyzer() *analysis.Analyzer {
	fc := floatcompare{}
	var flagSet flag.FlagSet

	flagSet.BoolVar(&fc.equalOnly, "equalOnly", false, "should the linter only search for == and !=")
	flagSet.BoolVar(&fc.skipTests, "skipTests", false, "should the linter execute on test files as well")

	return &analysis.Analyzer{
		Name:  "floatcompare",
		Doc:   "Search for float comparison, since these are potential errors",
		Run:   fc.run,
		Flags: flagSet,
	}
}

func (fc *floatcompare) isCheckExpr(node ast.Node, pass *analysis.Pass) {
	switch expr := node.(type) {
	case *ast.BinaryExpr:
		fc.checkBinExpr(expr, pass)
	case *ast.SwitchStmt:
		if fc.isFloat(expr.Tag, pass) {
			pass.Reportf(expr.Tag.Pos(), "float comparison with switch statement")
			return
		}
	}
}

func (fc *floatcompare) isFloat(expr ast.Expr, pass *analysis.Pass) bool {

	t := pass.TypesInfo.TypeOf(expr)
	if t == nil {
		return false
	}
	bt, ok := t.Underlying().(*types.Basic)
	if !ok {
		return false
	}
	if (bt.Info() & types.IsFloat) == 0 {
		return false
	}
	return true
}

func (fc *floatcompare) checkBinExpr(binExpr *ast.BinaryExpr, pass *analysis.Pass) {
	if fc.equalOnly && !(binExpr.Op == token.EQL || binExpr.Op == token.NEQ) {
		return
	}

	if !(binExpr.Op == token.EQL || binExpr.Op == token.LEQ || binExpr.Op == token.LSS || binExpr.Op == token.GEQ || binExpr.Op == token.GTR || binExpr.Op == token.NEQ) {
		return
	}

	if !fc.isFloat(binExpr.X, pass) || !fc.isFloat(binExpr.Y, pass) {
		return
	}

	pass.Reportf(binExpr.Pos(), "float comparison found %q",
		render(pass.Fset, binExpr))

}

func (fc *floatcompare) run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		if fc.skipTests && strings.HasSuffix(pass.Fset.Position(f.Pos()).Filename, "_test.go") {
			continue
		}

		ast.Inspect(f, func(node ast.Node) bool {
			fc.isCheckExpr(node, pass)
			return true
		})
	}
	return nil, nil
}
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
