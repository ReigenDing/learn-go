package main

import (
	"reflect"
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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5})
	want := []int{6, 9}
	if !reflect.DeepEqual(got, want) {

		t.Errorf("got %v want %v", got, want)
	}

}

func TestSumTails(t *testing.T) {
	t.Run("make sum of slice", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("salfly sum of empty slice", func(t *testing.T) {
		got := SumTails([]int{}, []int{3, 4})
		want := []int{0, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
