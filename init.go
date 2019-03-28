package main

import "fmt"

func main() {
	fmt.Print(double(3, 2))
}

func double(bullets, dragons int) bool {
	if bullets/dragons >= 2 {
		return true
	}
	return false
}
