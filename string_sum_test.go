package string_sum

import "testing"

func TestStringSum(t *testing.T) {
	//input := "\t\t\t   -20 -22       \t\t\t\t"
	input := " -3     +            4 "

	got, err := StringSum(input)
	want := "1"

	if got != want {
		t.Errorf("got %q want %q, with err %q", got, want, err)
	}
}
