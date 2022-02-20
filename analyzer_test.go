// Copyright: (c) 2022, Mathias Weber (mweb@gmx.ch)

package floatcompare

import (
	"os"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAllComparison_Run(t *testing.T) {
	a := NewAnalyzer()

	wd, _ := os.Getwd()
	analysistest.Run(t, wd+"/testdata", a, "all-comparison")
}

func TestEqualOnlyComparison_Run(t *testing.T) {
	a := NewAnalyzer()
	_ = a.Flags.Set("equalOnly", "true")

	wd, _ := os.Getwd()
	analysistest.Run(t, wd+"/testdata", a, "equal-only")
}

func TestSkipTestsComparison_Run(t *testing.T) {
	a := NewAnalyzer()
	_ = a.Flags.Set("skipTests", "true")

	wd, _ := os.Getwd()
	analysistest.Run(t, wd+"/testdata", a, "skipTests")
}

func TestWithTestsComparison_Run(t *testing.T) {
	a := NewAnalyzer()

	wd, _ := os.Getwd()
	analysistest.Run(t, wd+"/testdata", a, "withTests")
}
