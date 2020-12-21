package slicing

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	var tests = []struct {
		input         [][]int
		shouldBeError bool
	}{
		{[][]int{{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}}, false},
		{[][]int{{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9, 7},
			{7, 8, 9}}, true},
	}

	for i, test := range tests {
		_, err := NewMatrix(test.input)

		switch {
		case err == nil && test.shouldBeError:
			t.Errorf("In the test number %d should be returned error", i)

		case err != nil && !test.shouldBeError:
			t.Errorf("In the test number %d shouldn't be returned error", i)
		}

	}
}

func TestTransform(t *testing.T) {
	var tests = []struct {
		m           Matrix
		expectedOut []int
	}{
		{Matrix{
			size:   3,
			slices: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{Matrix{
			size:   4,
			slices: [][]int{{1, 2, 3, 1}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}},
		}, []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8}},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(test.expectedOut, test.m.Transform()) {
			t.Errorf("Test number %d crashed", i)
		}
	}
}
