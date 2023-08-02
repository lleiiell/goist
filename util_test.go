package goist

import "testing"

func TestContains(t *testing.T) {
	sli := []int{1, 2, 3}

	if !Contains(sli, 2) {
		t.Error("should Contains")
		return
	}

	if Contains(sli, 0) {
		t.Error("should Contains")
		return
	}

	if Contains(sli, "2") {
		t.Error("shouldn't Contains")
		return
	}
	if !Contains([]string{"a", "b", "c"}, "c") {
		t.Error("shouldn Contains")
		return
	}
}
