package main

import "fmt"

func main() {
	var n, a, b int64

	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(n * a * b * 2)
}
