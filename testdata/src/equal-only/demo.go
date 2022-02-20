// Copyright: (c) 2022, Mathias Weber (mweb@gmx.ch)

package main

func EqualCompareIfFloats() {
	x, y := 400., 500.
	if 300. == 100. { // want "float comparison found \"300. == 100.\""
		dummy()
	}
	if x == y { // want "float comparison found \"x == y\""
		dummy()
	}
	if 300.+200. == 10. { // want "float comparison found \"300.+200. == 10."
		dummy()
	}
	if 300 == 200 {
		dummy()
	}
}

func NotEqualCompareIfFloats() {
	x, y := 400., 500.
	if 300. != 100. { // want "float comparison found \"300. != 100.\""
		dummy()
	}
	if x != y { // want "float comparison found \"x != y\""
		dummy()
	}
}

func EqualCompareIfCustomType() {
	type number float64
	var x, y number = 300., 400.
	if x == y { // want "float comparison found \"x == y\""
		dummy()
	}
}

func EqualCompareIfFunctions() {
	if dummy() == dummy() { // want "float comparison found \"dummy.. == dummy..\""
		dummy()
	}
}

func EqualCompareIfNotSimpleType() {
	type demo struct {
		x float64
		y float64
	}

	k := demo{10., 20.}
	j := demo{22., 33.}

	if k == j {
		dummy()
	}
}

func GreaterLessCompareIfFloats() {
	if 300. >= 100. {
		dummy()
	}
	if 300. <= 100. {
		dummy()
	}
	if 300. < 100. {
		dummy()
	}
	if 300. > 100. {
		dummy()
	}
}

func SwitchStmtWithFloat() {
	switch 300. { // want "float comparison with switch statement"
	case 100.:
	}
}

func EqualCompareSwitchFloats() {
	switch {
	case 100. == 200.: // want "float comparison found \"100. == 200.\""
	}
}

func NotEqualCompareSwitchFloats() {
	switch {
	case 100. != 200.: // want "float comparison found \"100. != 200.\""
	}
}

func GreaterLessCompareSwitchFloats() {
	switch {
	case 100. <= 200.:
	case 100. < 200.:
	case 100. >= 200.:
	case 100. > 200.:
	}
}

func dummy() float64 { return 10. }
