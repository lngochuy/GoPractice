package tests

import (
	"testing"
	"tma/lngochuy/gopractice/largenum"
)

func TestMulSameSign(t *testing.T) {
	var a, b, r largenum.LargeNum

	a, _ = largenum.InitializeLargeNum("123")
	b, _ = largenum.InitializeLargeNum("3")
	r = a.Mul(b)
	if r.String() != "369" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("13")
	b, _ = largenum.InitializeLargeNum("85")
	r = a.Mul(b)
	if r.String() != "1105" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("123")
	b, _ = largenum.InitializeLargeNum("345")
	r = a.Mul(b)
	if r.String() != "42435" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("123.28")
	b, _ = largenum.InitializeLargeNum("345")
	r = a.Mul(b)
	if r.String() != "42531.6" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("-123.28")
	b, _ = largenum.InitializeLargeNum("-345.99")
	r = a.Mul(b)
	if r.String() != "42653.6472" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("123.28")
	b, _ = largenum.InitializeLargeNum("")
	r = a.Mul(b)
	if r.String() != "0" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("")
	b, _ = largenum.InitializeLargeNum("3.4")
	r = a.Mul(b)
	if r.String() != "0" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("")
	b, _ = largenum.InitializeLargeNum("")
	r = a.Mul(b)
	if r.String() != "0" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}
}
func TestMulDiffSign(t *testing.T) {
	var a, b, r largenum.LargeNum

	a, _ = largenum.InitializeLargeNum("123")
	b, _ = largenum.InitializeLargeNum("-3")
	r = a.Mul(b)
	if r.String() != "-369" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("-123")
	b, _ = largenum.InitializeLargeNum("345")
	r = a.Mul(b)
	if r.String() != "-42435" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("-123.28")
	b, _ = largenum.InitializeLargeNum("")
	r = a.Mul(b)
	if r.String() != "0" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}

	a, _ = largenum.InitializeLargeNum("")
	b, _ = largenum.InitializeLargeNum("-3.4")
	r = a.Mul(b)
	if r.String() != "0" {
		t.Logf("%s * %s = %s", a, b, r)
		t.Fail()
	}
}
