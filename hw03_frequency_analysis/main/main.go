package main

import (
	"bufio"
	"fmt"
	"os"

	hw03frequencyanalysis "github.com/MilkyWay-core/otus_golang/hw03_frequency_analysis"
)

func main() {
	var text string
	fmt.Println("Enter you text heare: ")
	consoleReader := bufio.NewReader(os.Stdin)
	text, _ = consoleReader.ReadString('\n')
	words := hw03frequencyanalysis.Top10(text)
	for _, word := range words {
		fmt.Println(word)
	}
}
