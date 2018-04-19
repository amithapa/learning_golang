package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("The square root of 4 is ", math.Sqrt(4))
}

func randNumber() {
	fmt.Println("Random number generated is: ", rand.Intn(100))
}

func main() {
	foo()
	randNumber()
}