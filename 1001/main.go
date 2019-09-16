package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []uint64

	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString(0)
	numbers, _ = parseInts(s)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%.4f\n", math.Sqrt(float64(numbers[len(numbers)-i-1])))
	}
}

func parseInts(s string) ([]uint64, error) {
	var (
		fields = strings.Fields(s)
		ints   = make([]uint64, len(fields))
		err    error
	)
	for i, f := range fields {
		ints[i], err = strconv.ParseUint(f, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}
