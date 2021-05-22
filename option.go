package niceID

// PrefixOption defines a generic option for the PrefixGenerator.
type PrefixOption func(*PrefixGenerator)

// WithPrefix modifies the default length of the prefix.
// The valid range is [0, n) where n should be a reasonable number.
// Very large n lead to a long computation.
func WithPrefix(length int) PrefixOption {
	return func(g *PrefixGenerator) {
		g.prefixLength = length
	}
}

// WithChecksum adds a checksum to every encoded id and validates the checksum
// upon decoding. If the checksum is not correct a zero id and an empty prefix
// are returned.
func WithChecksum() PrefixOption {
	return func(g *PrefixGenerator) {
		g.useChecksum = true
	}
}