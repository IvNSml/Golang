package main

import "fmt"

func main() {
	fmt.Print(SumEvenFibonacci(8), digitalRoot(256))
	// fmt.Print(digitalRoot(256))

}
func SumEvenFibonacci(limit int) (output int) {
	var current = 0
	var numbers = []int{1, 2}
	for current < limit {
		current = numbers[len(numbers)-2] + numbers[len(numbers)-1]
		numbers = append(numbers, current)
	}
	for _, value := range numbers {
		if value%2 == 0 {
			output += value
		}
	}
	return
}

func digitalRoot(n int) int {
	return (n-1)%9 + 1
}
