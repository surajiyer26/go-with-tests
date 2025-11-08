package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 4, 6}, []int{1, 3, 5, 7}, []int{0, 0})
	want := []int{12, 16, 0}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 4, 6}, []int{1, 3, 5, 7})
		want := []int{10, 15}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 3, 5, 7})
		want := []int{0, 15}
		checkSums(t, got, want)
	})
}
