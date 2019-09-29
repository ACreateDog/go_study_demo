package main

import "testing"

func TestFni(t *testing.T) {
	var (
		in = 7
		expected = 13
	)
	realRet := Fni(in)
	if realRet != expected {
		t.Errorf("TestFni,in=%+v,expected=%+v,realRet=%+v",in,expected,realRet)
	}
}