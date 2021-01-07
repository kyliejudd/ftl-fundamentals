package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		name string
	}
	testCases := []testCase{
		{name: "add two positives", a: 2, b: 2, want: 4},
		{name: "add two negatives", a: -2, b: -3, want: -5},
		{name: "add pos to neg", a: 4, b: -2, want: 2},
		{name: "add two fractions", a: 1.5, b: 2, want: 3.5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("The test: %v failed. Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		name string
	}
	testCases := []testCase{
		{name: "subtract two positives", a: 10, b: 2, want: 8},
		{name: "subtract two negatives", a: -5, b: -2, want: -3},
		{name: "subtract neg from a pos", a: -5, b: 2, want: -7},
		{name: "subtract fractions", a: 1.5, b: .5, want: 1.0},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("The test: %v failed, Subtract(%f, %f): want %v, got %v", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		name string
	}
	testCases := []testCase{
		{name: "multiply 2 positives", a: 10, b: 2, want: 20},
		{name: "multiply a pos by a neg", a: -5, b: 2, want: -10},
		{name: "multiply two negatives", a: -1, b: -5, want: 5},
		{name: "multiply fraction by whole", a: 1.5, b: 2, want: 3},
		{name: "multiply fraction by fraction", a: .5, b: .5, want: 0.25},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("the test: %v failed, Multiply(%f, %f): want %v, got %v", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{name: "divide by zero", a: 6, b: 0, want: 0, errExpected: true},
		{name: "divide two positives", a: 6, b: 2, want: 3, errExpected: false},
		{name: "One divided by two is one-half", a: 1, b: 2, want: 0.5, errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.want != got {
			t.Errorf("the test: %v failed, Divide(%f, %f): want %v, got %v", tc.name, tc.a, tc.b, tc.want, got)
		}
		if errReceived != tc.errExpected {
			t.Errorf("the test: %v failed, Divide(%f, %f): unexpected error code, want %v, got %v", tc.name, tc.a, tc.b, tc.errExpected, got)
		}
	}
}
