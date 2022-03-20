package main

import (
	"fmt"

	hw02unpackstring "github.com/MilkyWay-core/otus_golang/hw02_unpack_string"
)

func main() {
	var user_string string
	fmt.Println("Enter you text: ")
	num, err := fmt.Scan(&user_string)
	if (err != nil) || (num == 0) {
		fmt.Println(hw02unpackstring.ErrInvalidString)
		return
	}
	result, err := hw02unpackstring.Unpack(user_string)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
