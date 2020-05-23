package rotate

import (
	"reflect"
	"testing"
)

func TestRotateLeft1(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotateLeft(s, 2)
	expectedSlice := []int{2, 3, 4, 5, 0, 1}
	if !reflect.DeepEqual(s, expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, s)
	}
}

func TestRotateLeft2(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotateLeft(s, 5)
	expectedSlice := []int{5, 0, 1, 2, 3, 4}
	if !reflect.DeepEqual(s, expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, s)
	}
}

func TestRotateLeft3(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotateLeft(s, 7)
	expectedSlice := []int{1, 2, 3, 4, 5, 0}
	if !reflect.DeepEqual(s, expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, s)
	}
}

func TestRotateLeft4(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotateLeft(s, 0)
	expectedSlice := []int{0, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, s)
	}
}

func TestRotateRight1(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotateRight(s, 1)
	expectedSlice := []int{5, 0, 1, 2, 3, 4}
	if !reflect.DeepEqual(s, expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, s)
	}
}

func TestRotateRight2(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotateRight(s, 13)
	expectedSlice := []int{5, 0, 1, 2, 3, 4}
	if !reflect.DeepEqual(s, expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, s)
	}
}

func TestRotateRight3(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotateRight(s, 6)
	expectedSlice := []int{0, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, s)
	}
}
