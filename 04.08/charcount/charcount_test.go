package charcount

import (
	"bytes"
	"reflect"
	"testing"
)

func TestUpdateCounts(t *testing.T) {
	var tests = []struct {
		input rune
		want  CharacterCounter
	}{
		{input: 'a', want: CharacterCounter{"letters": 1}},
		{input: '1', want: CharacterCounter{"digits": 1}},
		{input: '🤨', want: CharacterCounter{"symbols": 1}},
	}
	for _, test := range tests {
		counts := make(CharacterCounter)
		updateCounts(counts, test.input)
		if !reflect.DeepEqual(counts, test.want) {
			t.Errorf("Expected %v, got %v for input %v", test.want, counts, test.input)
		}
	}
}

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input string
		want  CharacterCounter
	}{
		{input: "", want: CharacterCounter{}},
		{input: "abc日123", want: CharacterCounter{"letters": 4, "digits": 3}},
		{input: "abc日123🤨", want: CharacterCounter{"letters": 4, "digits": 3, "symbols": 1}},
	}
	for _, test := range tests {
		var stdin bytes.Buffer
		stdin.Write([]byte(test.input))
		counts := CharCount(&stdin)
		if !reflect.DeepEqual(counts, test.want) {
			t.Errorf("Expected %v, got %v for input %v", test.want, counts, test.input)
		}
	}
}
