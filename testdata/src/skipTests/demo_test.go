// Copyright: (c) 2022, Mathias Weber (mweb@gmx.ch)

package main

import "testing"

func Test_EqualCompareIfFloats(t *testing.T) {
	if 300. == 100. {
		dummy()
	}
}

func EqualCompareIfFloats(t *testing.T) {
	if 300. == 100. {
		dummy()
	}
}

func dummy() {}
