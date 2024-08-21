package main

import (
	"testing"
)

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expectedResult float32
	error bool
}{
	{name: "Divide 5 by 1", dividend: 5, divisor: 1, expectedResult: 5, error: false},
    {name: "Divide 0 by 1", dividend: 0, divisor: 1, expectedResult: 0, error: false},
    {name: "Divide 5 by 0", dividend: 5, divisor: 0, expectedResult: 0, error: true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests{
		got, err := Divide(tt.dividend, tt.divisor)
		if tt.error{
			if err == nil {
				t.Error("Expected error, got none")
			}
		}else {
			if err != nil {
				t.Error("Expected no error, got", err)
			}
		}

		if got != tt.expectedResult {
			t.Errorf("Expected %f but got %f", tt.expectedResult, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := Divide(5, 1)
	if err != nil {
		t.Error("Expected no error, got", err)
	}
}
