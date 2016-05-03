//If
// The if statement looks as it does in C or Java, except that the ( ) are gone and the { } are required.

package main

import "fmt"

func main() {
	sum := 1
	if sum == 0 {
		fmt.Println("Ok")	
	} else {
		fmt.Println("Not ok")
	}
}