// For
// Go has only one looping construct, the for loop.
// 
// The basic for loop looks as it does in C or Java, except that the ( ) are gone (they are not even optional) and the { } are required.
//

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}	
}
