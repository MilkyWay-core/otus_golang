package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

type wordCounted struct {
	count uint
	word  string
}

type countedList struct {
	words []wordCounted
}

func constructor(cl countedList) countedList {
	cl.words = make([]wordCounted, 0)
	return cl
}

func (cl countedList) add(wc wordCounted) countedList { // Was adding only uniq word
	for _, wordCounter := range cl.words {
		if wordCounter.word == wc.word {
			return cl
		}
	}
	cl.words = append(cl.words, wc)
	return cl
}

func Top10(str string) []string {
	fmt.Println("start")
	words := strings.Fields(str)
	cl := constructor(countedList{})
	for i, word := range words { // will counted  words in string
		var counterWord uint
		for n := i; n < len(words); n++ {
			if word == words[n] {
				counterWord++
			}
		}
		cl = cl.add(wordCounted{counterWord, word})
	}
	sort.SliceStable(cl.words, func(i, j int) bool { // sort in structure
		return cl.words[i].count > cl.words[j].count ||
			(cl.words[i].count == cl.words[j].count &&
				cl.words[i].word < cl.words[j].word)
	})
	fmt.Println(cl.words)
	resultArray := make([]string, 0)
	for i, word := range cl.words { // convert structure to string array
		if i > 9 {
			break
		}
		resultArray = append(resultArray, word.word)
	}

	return resultArray
}
