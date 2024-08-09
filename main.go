package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 1))
}

func Sum(value1, value2 int64) int64 {
	return value1 + value2
}
