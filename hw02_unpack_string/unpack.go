package hw02unpackstring

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var resultByte []byte
	for i := 0; i < len(str); i++ {
	swt:
		switch {
		case isNumber(str[i]) && i == 0: // first char must by dont number
			return "", ErrInvalidString
		case isNumber(str[i]) && isNumber(str[i-1]): // only number is not more 9
			return "", ErrInvalidString
		case isNumber(str[i]):
			s, _ := strconv.Atoi(string(str[i]))
			// if number is 0 then delete last char
			if s == 0 {
				resultByte = resultByte[:len(resultByte)-1]
				break swt
			}
			// copy char
			for n := 0; n < s-1; n++ {
				resultByte = append(resultByte, str[i-1])
			}
			break swt
		case isChar(str[i]):
			resultByte = append(resultByte, str[i])
			break swt
		default:
			return "", ErrInvalidString
		}
	}
	return string(resultByte), nil
}

func isNumber(char byte) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	return false
}

func isChar(char byte) bool {
	if char >= 65 && char <= 99 || char >= 97 && char <= 122 || char == 0x0A {
		return true
	}
	return false
}
