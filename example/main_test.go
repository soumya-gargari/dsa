package main

import (
	"testing"
)

func TestNegative(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{123, true},
		{121, false},
		{323, false},
		{-1, false},
		{3223, true},
	}
	for _, tc := range testCases {
		result := isPallindrome(tc.input)
		if result != tc.expected {
			t.Errorf("expected %v , got %v", tc.expected, result)
		}
	}

}
