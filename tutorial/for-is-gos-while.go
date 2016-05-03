// For is Go's "while"
// At that point you can drop the semicolons: C's while is spelled for in Go.

package main

import "fmt"

func main() {
	sum := 0
	for sum < 10 {
		fmt.Println(sum)
		sum++
	}
}