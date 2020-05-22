package reverseptr

import "testing"

func TestReverseOddLength(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	expectedArr := [5]int{5, 4, 3, 2, 1}
	reversePtrOddLength(&arr)
	if expectedArr != arr {
		t.Errorf("Expected %v, got %v", expectedArr, arr)
	}
}

func TestReverseEvenLength(t *testing.T) {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	expectedArr := [6]int{6, 5, 4, 3, 2, 1}
	reversePtrEvenLength(&arr)
	if expectedArr != arr {
		t.Errorf("Expected %v, got %v", expectedArr, arr)
	}
}
