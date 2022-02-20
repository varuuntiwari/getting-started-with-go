package main

import (
	"fmt"
)

func main() {
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Printf("Truncated number: %d\n", int64(input))
}