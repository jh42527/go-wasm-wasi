package main

import "fmt"

func main() {
	n := 100

	if n < 2 {
		fmt.Println(n)
	}

	var a, b int

	b = 1

	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	fmt.Println(b)
}
