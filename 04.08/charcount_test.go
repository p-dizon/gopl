package charcount

import (
	"reflect"
	"testing"
)

func TestUpdateCounts(t *testing.T) {
	var tests = []struct {
		input rune
		want  map[int]int
	}{
		{input: 'a', want: map[int]int{letters: 1}},
		{input: '1', want: map[int]int{digits: 1}},
		{input: '🤨', want: map[int]int{symbols: 1}},
	}
	for _, test := range tests {
		counts := make(map[int]int)
		updateCounts(counts, test.input)
		if !reflect.DeepEqual(counts, test.want) {
			t.Errorf("Expected %v, got %v for input %v", test.want, counts, test.input)
		}
	}
}

func TestCharCount(t *testing.T) {
	var tests = []struct {
		filename string
		want     map[int]int
	}{
		{filename: "assets/empty_file.txt", want: map[int]int{}},
		{filename: "assets/sample_text1.txt", want: map[int]int{letters: 4, digits: 3}},
		{filename: "assets/sample_text2.txt", want: map[int]int{letters: 4, digits: 3, symbols: 1}},
	}
	for _, test := range tests {
		counts := charCount(test.filename)
		if !reflect.DeepEqual(counts, test.want) {
			t.Errorf("Expected %v, got %v for filename %v", test.want, counts, test.filename)
		}
	}
}
