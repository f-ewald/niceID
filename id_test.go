package niceID

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefixGenerator_Generate(t *testing.T) {
	t.Parallel()

	g, err := NewPrefixGenerator( WithPrefix(10))
	assert.NoError(t, err)

	testCases := []struct{
		ID int64
	}{
		{0},
		{10},
		{100},
		{1000},
		{10000},
		{100000},
		{1000000},
		{10000000},
		{100000000},
		{1000000000},
		{10000000000},
		{100000000000},
		{1000000000000},
		{10000000000000},
	}

	for _, testCase := range testCases {
		s := g.Encode(testCase.ID)
		assert.True(t, len(s) >= 11)
	}
}

func BenchmarkPrefixGenerator_Generate(b *testing.B) {
	g, _ := NewPrefixGenerator( WithPrefix(10))
	for i := 0; i < b.N; i++ {
		g.Encode(int64(i))
	}
}

func TestRandomString(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		n int
	}{
		{1},
		{2},
		{3},
	}

	for _, testCase := range testCases {
		assert.Len(t, randomString(testCase.n), testCase.n)
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomString(10)
	}
}

func TestDecodeBase622(t *testing.T) {
	g, err := NewPrefixGenerator(WithPrefix(2))
	assert.NoError(t, err)
	id, s := g.Decode("aaKv")
	assert.Equal(t, int64(1337), id)
	assert.Equal(t, "aa", s)
}

func BenchmarkDecodeBase62(b *testing.B) {
	g, _ := NewPrefixGenerator(WithPrefix(2))
	const s = "aaLZ"
	for i := 0; i < b.N; i++ {
		g.Decode(s)
	}
}

func TestChecksum(t *testing.T) {
	t.Parallel()
	cs := checksum("test1")
	assert.Equal(t, int8(49), cs)
}

func BenchmarkChecksum(b *testing.B) {
	const s = "test"
	for i := 0; i < b.N; i++ {
		checksum(s)
	}
}
