package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct{
		in string
		expected string
	}{
		{
			in: " foo bar \n  hoge   fuga",
			expected: `{"column":["foo","bar"]}
{"column":["hoge","fuga"]}
`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.in, func(t *testing.T) {
			in := bytes.NewBufferString(tt.in)
			out := bytes.Buffer{}

			if err := run(in, &out); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expected, out.String())
		})
	}
}
