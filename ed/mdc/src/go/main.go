package main

import (
	"fmt"
)

func mdc(numA, numB int) int {
	if numA == 0 {
		return numB
	}

	if numB == 0 {
		return numA
	}

	return mdc(numB, numA % numB)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
