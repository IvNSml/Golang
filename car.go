package main

import "fmt"

func main() {

	fmt.Print(NbMounth(2000, 8000, 1000, 1.5))
}
func nbMounth(startOld, startNew, savings int, lossPercent float32) [2]int {
	var diff = startOld - startNew
	var n = diff / savings

}
