package niceID

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Generator encodes and decodes ids from a numeric representation to a string and back.
type Generator interface {
	Encode(int64) string
	Decode(string) (int64, string)
}

// PrefixGenerator
// SuffixGenerator
// FixedLengthGenerator

// PrefixGenerator encodes a number in a base 64 representation and adds a random prefix
// of a fixed size. Upon decoding, the prefix and the id are returned. A checksum validation
// is performed optionally.
//
// Example:
// gen := NewPrefixGenerator()
// gen.Encode(1337)
// // LZ
type PrefixGenerator struct {
	prefixLength int
	useChecksum  bool
}

func NewPrefixGenerator(options ...PrefixOption) (Generator, error) {
	g := &PrefixGenerator{
		prefixLength: 6,
		useChecksum:  false,
	}
	for _, opt := range options {
		opt(g)
	}
	return g, nil
}

func (pg *PrefixGenerator) Encode(id int64) (s string) {
	base62ID := EncodeBase64(id)
	s = randomString(pg.prefixLength) + base62ID

	// If a checksum should be used it will be appended.
	if pg.useChecksum {
		s += EncodeBase64(int64(checksum(s)))
	}
	return s
}

func (pg *PrefixGenerator) Decode(s string) (int64, string) {
	if len(s) <= pg.prefixLength {
		return 0, ""
	}
	if pg.useChecksum {
		if EncodeBase64(int64(checksum(s[:len(s)-1]))) != s[len(s):] {
			return 0, ""
		}
		s = s[:len(s)-1]
	}
	return DecodeBase62(s[pg.prefixLength:]), s[:pg.prefixLength]
}

// randomString creates a random string of length n using character from alphabet.
func randomString(n int) string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(alphabetCount))
		s[i] = string(alphabet[n.Int64()])
	}
	return strings.Join(s, "")
}

// checksum accepts a string and produces a numeric value between [0,62).
func checksum(s string) int8 {
	var sum int64
	for i := 0; i < len(s); i++ {
		sum += int64(s[i])
	}
	return int8(sum % alphabetCount)
}
