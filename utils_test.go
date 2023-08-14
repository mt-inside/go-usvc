package usvc

import (
	"testing"
)

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
