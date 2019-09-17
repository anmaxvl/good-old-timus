package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	switch {
	case n <= 4:
		fmt.Println("few")
		break
	case n <= 9:
		fmt.Println("several")
		break
	case n <= 19:
		fmt.Println("pack")
		break
	case n <= 49:
		fmt.Println("lots")
		break
	case n <= 99:
		fmt.Println("horde")
		break
	case n <= 249:
		fmt.Println("throng")
		break
	case n <= 499:
		fmt.Println("swarm")
		break
	case n <= 999:
		fmt.Println("zounds")
		break
	case n >= 1000:
		fmt.Println("legion")
		break
	}
}
