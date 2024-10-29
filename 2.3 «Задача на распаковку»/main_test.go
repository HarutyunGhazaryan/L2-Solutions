package main

import (
	"errors"
	"testing"
)

func TestUnpackStringWithBuilder(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
		expectedError  error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("incorrect string format")},
		{"", "", nil},
		{"qwe\\4\\5", "qwe45", nil},
		{"qwe\\45", "qwe44444", nil},
		{"qwe\\\\5", "qwe\\\\\\\\\\", nil},
	}

	for _, test := range tests {
		output, err := unpackStringWithBuilder(test.input)
		if output != test.expectedOutput || err != nil && err.Error() != test.expectedError.Error() {
			t.Errorf("For input %s, expected output %s and error %v, but got output %s and error %v",
				test.input, test.expectedOutput, test.expectedError, output, err,
			)
		}
	}
}

func TestUnpackStringWithConcat(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
		expectedError  error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("incorrect string format")},
		{"", "", nil},
		{"qwe\\4\\5", "qwe45", nil},
		{"qwe\\45", "qwe44444", nil},
		{"qwe\\\\5", "qwe\\\\\\\\\\", nil},
	}

	for _, test := range tests {
		output, err := unpackStringWithConcat(test.input)
		if output != test.expectedOutput || (err != nil && err.Error() != test.expectedError.Error()) {
			t.Errorf("For input '%s', expected output '%s' and error '%v', but got output '%s' and error '%v'",
				test.input, test.expectedOutput, test.expectedError, output, err)
		}
	}
}
