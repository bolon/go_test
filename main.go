package main

import "fmt"

func sum(a int, b int) bool {
	if a + b > 10 {
		return false
	}
	if a - b < 0 {
		return false
	}
	
	return true
}

func main() {
	fmt.Printf("Sum %v", sum(2, 3))
}