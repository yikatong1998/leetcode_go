package solution

import "testing"

func TestIntToRoman(t *testing.T) {
	input := 9
	output := "IX"
	if ans := intToRoman(input); ans != output {
		t.Errorf("intToRoman(%v) expected be %v, but %v got", input, output, ans)
	}
}
