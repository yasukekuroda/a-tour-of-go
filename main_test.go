package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{10, 20, 30},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
