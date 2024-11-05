package main

import (
	"fmt"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, -1, -2},
		{100, 200, 300},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d+%d", tt.a, tt.b), func(t *testing.T) {
			result := addNumbers(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("addNumbers(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
