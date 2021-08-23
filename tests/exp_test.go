package tests

import (
	"testing"
	"tma/lngochuy/gopractice/largenum"
)

func TestE(t *testing.T) {
	a, _ := largenum.InitializeLargeNum("1234.3")
	r := a.Exp(3)
	if r.String() != "1234300" {
		t.Logf("%s * 10e3 = %s", a, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1234.3")
	r = a.Exp(-3)
	if r.String() != "1.2343" {
		t.Logf("%s * 10e3 = %s", a, r)
		t.Fail()
	}
}
