package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// A CharacterCounter is a map that keeps track of
// unicode types in stdin
type CharacterCounter map[string]int

// CharCount counts the number of unicode types in stdin
func CharCount(stdin io.Reader) CharacterCounter {
	counts := make(CharacterCounter) // counts of Unicode character types
	in := bufio.NewReader(stdin)
	for {
		r, _, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		updateCounts(counts, r)
	}
	return counts
}

func updateCounts(c CharacterCounter, r rune) {
	switch {
	case unicode.IsLetter(r):
		c["letters"]++
	case unicode.IsDigit(r):
		c["digits"]++
	case unicode.IsSymbol(r):
		c["symbols"]++
	}
}
