package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("cllectio of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expect := 15
		if got != expect {
			t.Errorf("expected %d, but get %d", expect, got)
		}
	})

	t.Run("collections of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d give %v", got, want, numbers)
		}
	})

}
