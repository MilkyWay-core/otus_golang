package main

import (
	"fmt"

	hw02unpackstring "github.com/MilkyWay-core/otus_golang/hw02_unpack_string"
)

func main() {
	var userString string
	fmt.Println("Enter you text: ")
	num, err := fmt.Scan(&userString)
	if (err != nil) || (num == 0) {
		fmt.Println(hw02unpackstring.ErrInvalidString)
		return
	}
	result, err := hw02unpackstring.Unpack(userString)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
