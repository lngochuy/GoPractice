package tests

import (
	"testing"
	"tma/lngochuy/gopractice/largenum"
)

func TestAddSameSign(t *testing.T) {
	var a, b largenum.LargeNum

	a, _ = largenum.InitializeLargeNum("1234")
	b, _ = largenum.InitializeLargeNum("1234.")
	if r := a.Add(b); r.String() != "2468" {
		t.Logf("%s + %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("-1234")
	b, _ = largenum.InitializeLargeNum("-1234.")
	if r := a.Add(b); r.String() != "-2468" {
		t.Logf("%s + %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1235.")
	b, _ = largenum.InitializeLargeNum("1234.345")
	if r := a.Add(b); r.String() != "2469.345" {
		t.Logf("%s + %s = (%s)", a, b, r)
		t.Fail()
	}
}

func TestAddDiffSign(t *testing.T) {
	var a, b largenum.LargeNum

	a, _ = largenum.InitializeLargeNum("1234")
	b, _ = largenum.InitializeLargeNum("-1234.")
	if r := a.Add(b); r.String() != "0" {
		t.Logf("%s + %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("-1234")
	b, _ = largenum.InitializeLargeNum("1234.")
	if r := a.Add(b); r.String() != "0" {
		t.Logf("%s + %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1235.")
	b, _ = largenum.InitializeLargeNum("-1234.345")
	if r := a.Add(b); r.String() != ".655" {
		t.Logf("%s + %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1235.22222")
	b, _ = largenum.InitializeLargeNum("-1234.345")
	if r := a.Add(b); r.String() != ".87722" {
		t.Logf("%s + %s = (%s)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1230.22222")
	b, _ = largenum.InitializeLargeNum("-1234.345")
	if r := a.Add(b); r.String() != "-4.12278" {
		t.Logf("%s + %s = (%s)", a, b, r)
		t.Fail()
	}
}
