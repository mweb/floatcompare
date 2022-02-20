// Copyright: (c) 2022, Mathias Weber (mweb@gmx.ch)

package main

import (
	"github.com/mweb/floatcompare"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(floatcompare.NewAnalyzer())
}
