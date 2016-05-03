// Multiple results
// A function can return any number of results.
//
//The swap function returns two strings.

package main

import "fmt"

func swap(a, b int) (int, int) {
	a = a + b;
	b = a - b;
	a = a - b;
	return a,b;		
}
func main() {
	a := 10;
	b := 20
	fmt.Println("Initial A : ", a, " And B : ", b)
	a, b =  swap(a, b)
	fmt.Println("After A : ", a, " And B : ", b)
}