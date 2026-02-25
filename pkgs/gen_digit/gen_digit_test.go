package gendigit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateShortCode(t *testing.T) {
	length := 6
	res, err := GenerateShortCode("abcdefgh", length)
	require.NoError(t, err)
	require.Len(t, res, length)
}
