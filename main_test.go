package main

import "testing"

type testCase struct {
	input    int
	expected bool
}

func TestIsOdd(t *testing.T) {
	test := []testCase{
		{6, false},
		{5, true},
		{7, true},
		{8, false},
		{14, false},
		{97, true},
	}

	for _, tc := range test {
		result := IsOdd(tc.input)
		if result != tc.expected {
			t.Errorf("IsOdd(%d) = %v, expected %v", tc.input, result, tc.expected)
		}
	}
}
