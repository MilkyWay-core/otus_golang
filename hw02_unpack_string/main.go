package main

import (
	"fmt"

	"github.com/otus_golang/hw02_unpack_string/unpack"
)

func main() {
	var user_string string
	fmt.Println("Enter you text: ")
	num, err := fmt.Scan(&user_string)
	if (err != nil) || (num == 0) {
		fmt.Println(unpack.ErrInvalidString)
		return
	}
	result, err := unpack.Unpack(user_string)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
