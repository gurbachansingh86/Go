// Named return values
// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// 
// These names should be used to document the meaning of the return values.
// 
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// 
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.


package main

import "fmt"

func swap(a, b int) (x, y int) {
	a = a + b;
	b = a - b;
	a = a - b;
	x = a
	y = b
	return	
}
func main() {
	a := 10;
	b := 20
	fmt.Println("Initial A : ", a, " And B : ", b)
	a, b =  swap(a, b)
	fmt.Println("After A : ", a, " And B : ", b)
}
