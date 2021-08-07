package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	bytes := bytes.Buffer{}
	Greet(&bytes, "Chris")

	got := bytes.String()

	want := "hello, Chris"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
