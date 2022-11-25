package mathematics

import "testing"

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	expected := 15

	if expected != got {
		t.Errorf("expected %d got %d", expected, got)
	}
}