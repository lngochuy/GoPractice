package tests

import (
	"testing"
	"tma/lngochuy/gopractice/largenum"
)

func TestSubSameSign(t *testing.T) {
	var a, b largenum.LargeNum

	a, _ = largenum.InitializeLargeNum("1235.")
	b, _ = largenum.InitializeLargeNum("1234.345")
	if r := a.Sub(b); r.String() != ".655" {
		t.Logf("%s - %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1235.")
	b, _ = largenum.InitializeLargeNum("1236.345")
	if r := a.Sub(b); r.String() != "-1.345" {
		t.Logf(r.Debug())
		t.Logf("%s - %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1235.")
	b, _ = largenum.InitializeLargeNum("1235")
	if r := a.Sub(b); r.String() != "0" {
		t.Logf("%s - %s = (%s)", a, b, r)
		t.Fail()
	}
}

func TestSubDiffSign(t *testing.T) {
	var a, b largenum.LargeNum

	a, _ = largenum.InitializeLargeNum("1235.")
	b, _ = largenum.InitializeLargeNum("-1234.345")
	if r := a.Sub(b); r.String() != "2469.345" {
		t.Logf("%s - %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("-1235.")
	b, _ = largenum.InitializeLargeNum("1236.345")
	if r := a.Sub(b); r.String() != "-2471.345" {
		t.Logf("%s - %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("-1235.")
	b, _ = largenum.InitializeLargeNum("1235")
	if r := a.Sub(b); r.String() != "-2470" {
		t.Logf("%s - %s = (%s)", a, b, r)
		t.Fail()
	}
}
