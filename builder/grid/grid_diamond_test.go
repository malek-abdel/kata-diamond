package grid

import "testing"

func TestPrintLetterC(t *testing.T) {
	gridDiamondBuilder := GridDiamondBuilder{}
	got, _ := gridDiamondBuilder.Print("C")
	expected := "  A  \n B B \nC   C\n B B \n  A  "
	if got != expected {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}

func TestPrintLetterA(t *testing.T) {
	gridDiamondBuilder := GridDiamondBuilder{}
	got, _ := gridDiamondBuilder.Print("C")
	expected := "A"
	if got != expected {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}

func TestPrintNotALetter(t *testing.T) {
	gridDiamondBuilder := GridDiamondBuilder{}
	got, err := gridDiamondBuilder.Print("8")
	expected, expected_err := "", "You should pass a capital letter between A and Z"
	if got != expected {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
	if err.Error() != expected_err {
		t.Errorf("got: %q, expected: %q", err.Error(), expected_err)
	}
}
