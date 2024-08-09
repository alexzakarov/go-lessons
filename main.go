package main

import "fmt"

func main() {
	name := "Hakan"
	surname := "Yenidogan"

	fmt.Println(Sum(1, 1))
	fmt.Println(msg(name, surname))
	fmt.Println(sum(5, 10, 15, 31))
}

func Sum(value1, value2 int64) int64 {
	return value1 + (value2 * 2)
}

var numbers = []int{1, 2, 3, 4, 5, 6, 7}

func sum(number ...int) int {
	sum := 0
	for _, value := range number {
		sum += value
	}
	return sum
}

func msg(name, surname string) string {
	return fmt.Sprintf("Hello %s, %s", name, surname)
}
