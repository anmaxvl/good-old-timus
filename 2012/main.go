package main

import "fmt"

func main() {
	var f int

	fmt.Scan(&f)

	if f < 7 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
