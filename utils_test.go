package utils

import (
	"testing"
)

func Test_EqualFloat64(t *testing.T) {
	one, two, three, dec := 1.11111111, 1.11111111, 1.111111111, 0.0000000001

	if EqualFloaf64(one, two, dec) == false {
		t.Fail()
	}

	if EqualFloaf64(one, three, dec) == true {
		t.Fail()
	}

}
