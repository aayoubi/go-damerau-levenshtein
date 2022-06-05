package distance

import "testing"

func TestDistances(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected int
	}{
		{"", "", 0},
		{"a", "", 1},
		{"", "a", 1},
		{"aaa", "aab", 1},
		{"aba", "acb", 2},
		{"golang", "java", 5},
	}
	for _, test := range tests {
		if res := DistanceDamerauLevenshtein(test.s1, test.s2); res != test.expected {
			t.Errorf("Error computing distance between [%s] and [%s], got: %d, want: %d", test.s1, test.s2, res, test.expected)
		}
	}
}
