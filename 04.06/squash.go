package squash

import (
	"unicode"
	"unicode/utf8"
)

func squashUnicodeSpaces(b []byte) []byte {
	// squashedSlice will overwrite values in slice b
	// starting with b[0]
	squashedSlice := b[:0]
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if !unicode.IsSpace(r) {
			squashedSlice = append(squashedSlice, b[i:i+size]...)
		} else if isInitialSpaceChar(squashedSlice) {
			squashedSlice = append(squashedSlice, ' ')
		}
		i += size
	}
	return squashedSlice
}

func isInitialSpaceChar(b []byte) bool {
	return (len(b) == 0) || (b[len(b)-1] != ' ')
}
