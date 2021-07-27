package main

import "testing"

func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},

		{"1", 1},
		{"b", 1},
		{"bbbbbbbb", 1},
	}

	for _, tt := range tests {
		actual := subStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	tests := []struct {
		s   string
		ans int
	}{
		{
			"bbbbbbbbbbbbbbbbbbbbbb",
			1,
		},
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			actual := subStr(tt.s)
			if actual != tt.ans {
				b.Errorf("got %d for input %s expected %d", actual, tt.s, tt.ans)
			}
		}
	}
}
