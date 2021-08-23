package tests

import (
	"testing"
	"tma/lngochuy/gopractice/largenum"
)

func TestDivSameSign(t *testing.T) {
	var a, b, r largenum.LargeNum

	a, _ = largenum.InitializeLargeNum("1234.9876")
	b, _ = largenum.InitializeLargeNum("13")
	r = a.Div(b, -11)
	if r.String() != "94.99904615384" {
		t.Logf("%s / %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("1234.9876")
	b, _ = largenum.InitializeLargeNum("13")
	r = a.Div(b, 1)
	if r.String() != "90" {
		t.Logf("%s / %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("9160514.97812")
	b, _ = largenum.InitializeLargeNum("892.63500")
	r = a.Div(b, -7)
	if r.String() != "10262.3300432" {
		t.Logf("%s / %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("")
	b, _ = largenum.InitializeLargeNum("892.63500")
	r = a.Div(b, -7)
	if r.String() != "0" {
		t.Logf("%s / %s = %s", a, b, r)
		t.Fail()
	}
}


func TestDivDiffSign(t *testing.T) {
	var a, b, r largenum.LargeNum

	a, _ = largenum.InitializeLargeNum("1234.9876")
	b, _ = largenum.InitializeLargeNum("-13")
	r = a.Div(b, -11)
	if r.String() != "-94.99904615384" {
		t.Logf("%s / %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("-1234.9876")
	b, _ = largenum.InitializeLargeNum("13")
	r = a.Div(b, 1)
	if r.String() != "-90" {
		t.Logf("%s / %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("-9160514.97812")
	b, _ = largenum.InitializeLargeNum("892.63500")
	r = a.Div(b, -7)
	if r.String() != "-10262.3300432" {
		t.Logf("%s / %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("")
	b, _ = largenum.InitializeLargeNum("-892.63500")
	r = a.Div(b, -7)
	if r.String() != "0" {
		t.Logf("%s / %s = %s", a, b, r)
		t.Fail()
	}
}

