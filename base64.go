package niceID

import (
	"math"
	"strings"
)

// alphabet as defined in https://en.wikipedia.org/wiki/Base62
// with an additional two url-safe characters
const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

var alphabetCount = int64(len(alphabet))

func EncodeBase64(id int64) string {
	if id <= 0 {
		return "0"
	}
	chars := make([]string, 0)
	for id > 0 {
		n := id % alphabetCount
		chars = append([]string{string(alphabet[n])}, chars...)
		id /= alphabetCount
	}
	return strings.Join(chars, "")
}

func DecodeBase62(s string) int64 {
	if s == "" {
		return 0
	}
	var sum int64
	split := strings.Split(s, "")
	for i := len(split)-1; i >= 0; i-- {
		idx := strings.Index(alphabet, split[i])
		sum += int64(idx) * int64(math.Pow(float64(alphabetCount), float64(len(s) -1 - i)))
	}
	return sum
}
