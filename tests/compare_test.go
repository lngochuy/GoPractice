package tests

import (
	"testing"
	"tma/lngochuy/gopractice/largenum"
)

func TestCompare(t *testing.T) {
	var a, b largenum.LargeNum

	a, _ = largenum.InitializeLargeNum("1234")
	b, _ = largenum.InitializeLargeNum("4567")
	if r := a.Compare(b); r != -1 {
		t.Logf("%s < %s fails (%d)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1234")
	b, _ = largenum.InitializeLargeNum("1234")
	if r := a.Compare(b); r != 0 {
		t.Logf("%s == %s fails (%d)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1234.56")
	b, _ = largenum.InitializeLargeNum("1234")
	if r := a.Compare(b); r != 1 {
		t.Logf("%s > %s fails (%d)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("6.789")
	b, _ = largenum.InitializeLargeNum("4567")
	if r := a.Compare(b); r != -1 {
		t.Logf("%s < %s fails (%d)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1212.1212")
	b, _ = largenum.InitializeLargeNum("1212.1213")
	if r := a.Compare(b); r != -1 {
		t.Logf("%s < %s fails (%d)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum(".234")
	b, _ = largenum.InitializeLargeNum("0.234")
	if r := a.Compare(b); r != 0 {
		t.Logf("%s == %s fails (%d)", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1234.000")
	b, _ = largenum.InitializeLargeNum("1234.")
	if r := a.Compare(b); r != 0 {
		t.Logf("%s == %s fails (%d)", a, b, r)
		t.Fail()
	}
}
