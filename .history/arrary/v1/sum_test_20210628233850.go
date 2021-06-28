package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	expect := 15
	if got != expect {
		t.Errorf("expected %d, but get %d", expect, got)
	}

}
