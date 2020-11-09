package main

func sum(a int, b int) bool {
	if a + b > 5 {
		return false
	}
	if a - b < 0 {
		return false
	}
	
	return true
}