package main

import (
	"fmt"	
)

// Equivelent to func add(x float64, y float64)
func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a,b
}

func main() {
	/*

	Equivalent to
	var num1 float64 = 5.7
	var num2 float64 = 9.7
	They get unpacked
	
	var num1, num2 float64 = 5.7, 9.7
	We cannot change the type of the data it can store the Go will come once and after it is compiled we cannot change the data type
	

	num1, num2 := 5.7, 9.7

	fmt.Println("The sum of a and b is ", add(num1, num2))

	var a int = 64
	var b float64 = float64(a)
	// Here x will be int
	x := a
	*/

	w1, w2 := "bunny", "shonuu"
	// fmt.Println(w1, w2)	
	fmt.Println(multiple(w1,w2))
}