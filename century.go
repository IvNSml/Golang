package main

import "fmt"

func main() {

	fmt.Print(century(2001))
}

func century(year int) int {
	return (year-1)/100 + 1
}
