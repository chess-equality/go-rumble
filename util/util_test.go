package util

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAppendSlice(t *testing.T) {

	byteSlice := []byte{'G', 'o', 'l', 'a', 'n', 'g'}

	assert := assert.New(t)

	var tests = []struct {
		input    []byte
		appended []byte
		expected []byte
	}{
		{byteSlice, []byte{' ', 'F', 'T', 'W', '!'}, []byte{'G', 'o', 'l', 'a', 'n', 'g', ' ', 'F', 'T', 'W', '!'}},
	}

	for _, test := range tests {

		// Test
		actual := Append(test.input, test.appended)
		log.Printf(">>>> actual = %s", string(actual))

		assert.Equal(actual, test.expected)
	}
}
