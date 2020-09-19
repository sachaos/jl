package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	testCases := []struct{
		in string
		expected []string
	}{
		{
			in: " foo bar ",
			expected: []string{"foo", "bar"},
		},
		{
			in: " foo  bar ",
			expected: []string{"foo", "bar"},
		},
		{
			in: " foo   bar ",
			expected: []string{"foo", "bar"},
		},
		{
			in: " foo    bar ",
			expected: []string{"foo", "bar"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.in, func(t *testing.T) {
			assert.Equal(t, tt.expected, Split(tt.in))
		})
	}
}
