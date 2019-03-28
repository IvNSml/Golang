package main

import "math"

func main() {
	var test_unit []float64 = []float64{}
	print(Gps(15, test_unit))
}

func Gps(s int, x []float64) int {
	if s == 0 {
		panic("ERROR: Division by zero")
	}
	if len(x) == 0 {
		return 0
	}
	var speed []float64
	for index := range x[:len(x)-1] {
		speed = append(speed, ((x[index+1]-x[index])*3600)/float64(s))
	}
	var max float64 = 0
	for _, value := range speed {
		if max < value {
			max = value
		}
	}
	if max <= 1 {
		return 0
	}
	return int(math.Floor(max))
}

func DigitalRoot(n int) int {
	for true {
		if n < 10 {
			return n
		}
		n = Endless(n)
	}
	return -1
}
func Endless(number int) int {
	if number < 10 {
		return number
	}
	return number%10 + Endless(number/10)
}
