package squash

import (
	"reflect"
	"testing"
)

func TestSquash1(t *testing.T) {
	input := []byte(" asdf\n\t \nqwerty")
	expectedOutput := []byte(" asdf qwerty")
	output := squashUnicodeSpaces(input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}

func TestSquash2(t *testing.T) {
	input := []byte("\t\t\t\t\nqwerty")
	expectedOutput := []byte(" qwerty")
	output := squashUnicodeSpaces(input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}

func TestSquash3(t *testing.T) {
	input := []byte("qwerty\t\t\t\t\n")
	expectedOutput := []byte("qwerty ")
	output := squashUnicodeSpaces(input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}

func TestSquash4(t *testing.T) {
	input := []byte("a\ts\td\tf")
	expectedOutput := []byte("a s d f")
	output := squashUnicodeSpaces(input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}

func TestSquashEmptySlice(t *testing.T) {
	input := []byte("")
	expectedOutput := []byte("")
	output := squashUnicodeSpaces(input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}
