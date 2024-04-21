package main

import (
	"testing"
)

func TestIsPallindrome(t *testing.T) {
	testCases := []struct {
		data     int
		expected bool
	}{
		{121, true},
		{342, false},
		{989, true},
	}
	for _, tc := range testCases {
		result := isPallindrome(tc.data)
		if result != tc.expected {
			t.Errorf("expected %v, go %v", tc.expected, result)
		}
	}
}
