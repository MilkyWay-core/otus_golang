package unpack

import (
	"errors"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "aaA2b", expected: "aaAAb"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "фыва10", "as2 fds"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
func TestUnpackRandString(t *testing.T) {
	const goodChar = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321"
	var resultByts []byte
	n := 0
	for n < rand.Intn(100) {
		randChar := goodChar[rand.Intn(len(goodChar))]
		if isNumber(randChar) && (n != 0 || !isNumber(resultByts[n-1])) {
			resultByts = append(resultByts, goodChar[rand.Intn(len(goodChar))])
			n++
		}
		if isChar(randChar) {
			resultByts = append(resultByts, goodChar[rand.Intn(len(goodChar))])
			n++
		}
	}
	resultString := string(resultByts)
	t.Run(resultString, func(t *testing.T) {
		_, err := Unpack(resultString)
		require.NoError(t, err)
	})

}
