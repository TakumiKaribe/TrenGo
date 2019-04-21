package service

import (
	"testing"
)

func TestMin(t *testing.T) {
	cases := []struct {
		expected int
		lhs      int
		rhs      int
	}{
		{expected: 0, lhs: 0, rhs: 1},
		{expected: -1, lhs: 0, rhs: -1},
		{expected: 100, lhs: 100, rhs: 101},
	}

	for _, c := range cases {
		if c.expected != Min(c.lhs, c.rhs) {
			t.Errorf("Min is not %d, got= %d", c.expected, Min(c.lhs, c.rhs))
		}
	}
}
