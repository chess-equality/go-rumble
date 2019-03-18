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
		append   []byte
		expected []byte
	}{
		{byteSlice, []byte{' ', 'F', 'T', 'W', '!'}, []byte{'G', 'o', 'l', 'a', 'n', 'g', ' ', 'F', 'T', 'W', '!'}},
	}

	for _, test := range tests {

		actual := Append(test.input, test.append)
		log.Printf(">>>> actual = %s", string(actual))

		assert.Equal(Append(test.input, test.append), test.expected)
	}
}
