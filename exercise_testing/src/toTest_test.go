package totest

import (
	"testing"
)

func Test_testmeDatadriver(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{2, 4},
		{3, 9},
		{4, 16},
		{1, 1},
		{0, 0}
	}

	
}
func Test_testme(t *testing.T) {
	input := 2
	expected := 5
	result := testme(input)
	if expected != result {
		t.Fail()
	}
}
