package squash

import "unicode"

// func squashUnicodeSpaces(b []byte) []byte {
// 	squashedSlice := b[:1]
// 	for _, char := range b[1:] {
// 		if !unicode.IsSpace(rune(char)) {
// 			squashedSlice = append(squashedSlice, char)
// 		} else if squashedSlice[len(squashedSlice)-1] != ' ' {
// 			squashedSlice = append(squashedSlice, ' ')
// 		}
// 	}
// 	return squashedSlice
// }

func squashUnicodeSpaces(b []byte) []byte {
	// squashedSlice will overwrite values in slice b
	squashedSlice := b[:0]
	for _, char := range b[:] {
		if !unicode.IsSpace(rune(char)) {
			squashedSlice = append(squashedSlice, char)
		} else if isInitialSpaceChar(squashedSlice) {
			squashedSlice = append(squashedSlice, ' ')
		}
	}
	return squashedSlice
}

func isInitialSpaceChar(b []byte) bool {
	return (len(b) == 0) || (b[len(b)-1] != ' ')
}
