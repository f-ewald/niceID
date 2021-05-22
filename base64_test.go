package niceID

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeBase62(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		id int64
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{9, "9"},
		{10, "A"},
		{11, "B"},
		{61, "z"},
		{40346747266, "base62"},
		{1337, "Kv"},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, EncodeBase64(testCase.id))
	}
}

func BenchmarkEncodeBase62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeBase64(int64(i))
	}
}


func TestDecodeBase62(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		id string
		expected int64
	}{
		{"0", 0},
		{"1", 1},
		{"9", 9},
		{"A", 10},
		{"B", 11},
		{"z", 61},
		{"Fe", 1000},
		{"base62", 40346747266},
		{"Kv", 1337},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, DecodeBase62(testCase.id))
	}
}
