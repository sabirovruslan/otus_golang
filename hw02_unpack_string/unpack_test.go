package hw02unpackstring

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "", expected: ""},
		{input: " ", expected: " "},
		{input: "a", expected: "a"},
		{input: "aa", expected: "aa"},
		{input: "abc", expected: "abc"},
		{input: "a0", expected: ""},
		{input: "a0", expected: ""},
		{input: "a1", expected: "a"},
		{input: "a2", expected: "aa"},
		{input: "a3", expected: "aaa"},
		{input: "a-1", expected: "a-"},
		{input: "a-0", expected: "a"},
		{input: " 0", expected: ""},
		{input: " 1", expected: " "},
		{input: " 2", expected: "  "},
		{input: "a2a", expected: "aaa"},
		{input: "a3a", expected: "aaaa"},
		{input: "a3c®", expected: "aaac®"},
		{input: "abc4a2", expected: "abccccaa"},
		{input: "a0b0c0d0", expected: ""},
		{input: "a w3 a", expected: "a www a"},
	}

	for _, tc := range tests {
		tc = tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	tests := []string{"0", "1", "2q", "01", "a01"}
	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
