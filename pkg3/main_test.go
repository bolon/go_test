package pkg3

import "testing"

func TestSum(t *testing.T) {
	var expected = true
	var actual = sum(3, 2)
	if expected != actual {
		t.Errorf("%v == %v\n", expected, actual)
	}
  
	expected = false
  actual = sum(6, 10)
  if expected != actual {
    t.Errorf("%v == %v\n", expected, actual)
  }
}