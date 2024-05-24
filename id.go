package id

import (
	cryptoRand "crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

// Generate returns a random string of the specified length using the provided alphabet.
func Generate(alphabet string, length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0, got %d", length)
	}
	alphabetSize := big.NewInt(int64(len(alphabet)))
	if alphabetSize.Int64() == 0 {
		return "", errors.New("alphabet must not be empty")
	}

	bytes := make([]byte, length)
	for i := range bytes {
		n, err := cryptoRand.Int(cryptoRand.Reader, alphabetSize)
		if err != nil {
			return "", fmt.Errorf("generating random number: %w", err)
		}
		bytes[i] = alphabet[n.Int64()]
	}

	return string(bytes), nil
}

// Generator returns a function that generates a random string of the specified length
// using the provided alphabet.
func Generator(alphabet string, length int) func() (string, error) {
	return func() (string, error) {
		return Generate(alphabet, length)
	}
}
