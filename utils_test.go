package usvc

import (
	"testing"
)

func TestMin(t *testing.T) {
	cases := []struct {
		x        int
		y        int
		expected int
	}{
		{3, 1, 1},
		{5, 1, 1},
		{1, 5, 1},
		{1, 1, 1},
		{-1, 1, -1},
		{1, -1, -1},
		{0, 1, 0},
		{0, -1, -1},
	}

	for _, cse := range cases {
		got := Min(cse.x, cse.y)
		if got != cse.expected {
			t.Errorf("Min(%d, %d): got %d, expected %d.", cse.x, cse.y, got, cse.expected)
		}
	}
}
func TestMax(t *testing.T) {
	cases := []struct {
		x        int
		y        int
		expected int
	}{
		{3, 1, 3},
		{5, 1, 5},
		{1, 5, 5},
		{1, 1, 1},
		{-1, 1, 1},
		{1, -1, 1},
		{0, 1, 1},
		{0, -1, 0},
	}

	for _, cse := range cases {
		got := Max(cse.x, cse.y)
		if got != cse.expected {
			t.Errorf("Max(%d, %d): got %d, expected %d.", cse.x, cse.y, got, cse.expected)
		}
	}
}

func TestTernaryInt(t *testing.T) {
	cases := []struct {
		test     bool
		x        int
		y        int
		expected int
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
	}

	for _, cse := range cases {
		got := Ternary(cse.test, cse.x, cse.y)
		if got != cse.expected {
			t.Errorf("(%t ? %v : %v): got %v, expected %v.", cse.test, cse.x, cse.y, got, cse.expected)
		}
	}
}
func TestTernaryString(t *testing.T) {
	cases := []struct {
		test     bool
		x        string
		y        string
		expected string
	}{
		{true, "a", "b", "a"},
		{false, "a", "b", "b"},
	}

	for _, cse := range cases {
		got := Ternary(cse.test, cse.x, cse.y)
		if got != cse.expected {
			t.Errorf("(%t ? %v : %v): got %v, expected %v.", cse.test, cse.x, cse.y, got, cse.expected)
		}
	}
}
