package charcount

import (
	"io/ioutil"
	"log"
	"unicode"
	"unicode/utf8"
)

const (
	letters = iota
	digits
	symbols
)

func charCount(filename string) map[int]int {
	counts := make(map[int]int)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(contents); {
		r, size := utf8.DecodeRune(contents[i:])
		updateCounts(counts, r)
		i += size
	}
	return counts
}

func updateCounts(c map[int]int, r rune) {
	switch {
	case unicode.IsLetter(r):
		c[letters]++
	case unicode.IsDigit(r):
		c[digits]++
	case unicode.IsSymbol(r):
		c[symbols]++
	}
}
