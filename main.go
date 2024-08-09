package main

import "fmt"

func main() {
	name := "Hakan"
	surname := "Yenidogan"

	fmt.Println(Sum(1, 1))
	fmt.Println(msg(name, surname))
}

func Sum(value1, value2 int64) int64 {
	return value1 + value2
}

func msg(name, surname string) string {
	return fmt.Sprintf("Hello %s, %s", name, surname)
}
