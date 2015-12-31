// Functions continued
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
//
// In this example, we shortened
//
// x int, y int
// to
//x, y int

package main

import (
	"fmt"
)

func calc(a, b int) (int) {
	return a + b
}

func main() {
	fmt.Println("Sum of two number " ,calc(2, 3))
}