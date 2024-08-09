package main

import "fmt"

func main() {
	fmt.Println(Sum(1.0, 1.0))
}

func Sum(value1, value2 float64) float64 {
	return value1 + value2
}
