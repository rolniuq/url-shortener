package gendigit

import (
	"crypto/rand"
	"math/big"
)

func GenerateShortCode(charset string, length int) (string, error) {
	result := make([]byte, length)

	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", nil
		}

		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}
