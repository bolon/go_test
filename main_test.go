package main

import "testing"

func TestSum(t *testing.T) {
	var expected = true
	var actual = sum(2, 3)
	if expected != actual {
		t.Logf("%v == %v\n", expected, actual)
	}
}