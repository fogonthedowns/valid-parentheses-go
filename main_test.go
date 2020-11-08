package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	in := "[[[{{{(())}}}]]]"
	out := isValid(in)
	want := true
	if out != want {
		t.Errorf("got %t, want %t", out, want)
	}
}
