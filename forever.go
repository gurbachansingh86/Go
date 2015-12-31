// Forever
// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

package main

import "fmt"

func main() {
	sum := 0
	for {
		fmt.Println(sum)
		sum++
	}
}