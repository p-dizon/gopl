package reverse

import (
	"reflect"
	"testing"
)

func TestReverseUtf8(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("abcd日"), []byte("日dcba")},
		{[]byte("日abcd"), []byte("dcba日")},
		{[]byte("日abcd日"), []byte("日dcba日")},
	}
	for _, test := range tests {
		output := reverseUtf8Slice(test.input)
		if !reflect.DeepEqual(output, test.want) {
			t.Errorf("Wanted %v, got %v", test.want, output)
		}
	}
}

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []byte("abcd日")
		reverseUtf8Slice(input)
	}
}
