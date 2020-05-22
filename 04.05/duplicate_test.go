package duplicate

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicate1(t *testing.T) {
	s := []string{"hello", "hello", "there"}
	expectedSplice := []string{"hello", "there"}
	s = removeAdjacentDuplicates(s)
	if !reflect.DeepEqual(s, expectedSplice) {
		t.Errorf("Expected %v, got %v", expectedSplice, s)
	}
}

func TestRemoveDuplicate2(t *testing.T) {
	s := []string{"hello", "hello", "how", "how", "there"}
	expectedSplice := []string{"hello", "how", "there"}
	s = removeAdjacentDuplicates(s)
	if !reflect.DeepEqual(s, expectedSplice) {
		t.Errorf("Expected %v, got %v", expectedSplice, s)
	}
}

func TestRemoveDuplicate3(t *testing.T) {
	s := []string{"hello", "how", "there"}
	expectedSplice := []string{"hello", "how", "there"}
	s = removeAdjacentDuplicates(s)
	if !reflect.DeepEqual(s, expectedSplice) {
		t.Errorf("Expected %v, got %v", expectedSplice, s)
	}
}

func TestRemoveDuplicate4(t *testing.T) {
	s := []string{"hello"}
	expectedSplice := []string{"hello"}
	s = removeAdjacentDuplicates(s)
	if !reflect.DeepEqual(s, expectedSplice) {
		t.Errorf("Expected %v, got %v", expectedSplice, s)
	}
}

func TestRemoveDuplicate5(t *testing.T) {
	s := []string{"hello", "there"}
	expectedSplice := []string{"hello", "there"}
	s = removeAdjacentDuplicates(s)
	if !reflect.DeepEqual(s, expectedSplice) {
		t.Errorf("Expected %v, got %v", expectedSplice, s)
	}
}

func TestRemoveDuplicate6(t *testing.T) {
	s := []string{"hello", "hello"}
	expectedSplice := []string{"hello"}
	s = removeAdjacentDuplicates(s)
	if !reflect.DeepEqual(s, expectedSplice) {
		t.Errorf("Expected %v, got %v", expectedSplice, s)
	}
}

func TestRemoveDuplicate7(t *testing.T) {
	s := []string{"hello", "hello", "how", "how", "there", "there"}
	expectedSplice := []string{"hello", "how", "there"}
	s = removeAdjacentDuplicates(s)
	if !reflect.DeepEqual(s, expectedSplice) {
		t.Errorf("Expected %v, got %v", expectedSplice, s)
	}
}
