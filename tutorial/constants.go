// Constants
// 
// Constants are declared like variables, but with the const keyword.
// 
// Constants can be character, string, boolean, or numeric values.
// 
// Constants cannot be declared using the := syntax.

package main

import "fmt"

const (
 	end = "End Of World"
 	world = "Hello I am here with "
)

func main() {
	const world = "Call Me"
	fmt.Println(world, end);
}