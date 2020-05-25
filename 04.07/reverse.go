package reverse

import "unicode/utf8"

func reverseUtf8Slice(b []byte) []byte {
	reversedSlice := []byte{}
	for len(b) > 0 {
		_, size := utf8.DecodeLastRune(b)
		indexOfLastRune := len(b) - size
		reversedSlice = append(reversedSlice, b[indexOfLastRune:]...)
		b = b[:indexOfLastRune]
	}
	return reversedSlice
}
