package reverse_integer

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if ans := reverse2(-2147483412); ans != -2143847412 {
		t.Errorf("reverse(123) expected be 321, but %d got", ans)
	}
}
