// Exported names
// In Go, a name is exported if it begins with a capital letter.
// 
// When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
// 
// Foo is an exported name, as is FOO. The name foo is not exported.
// 
// Run the code. Notice the error message.
// 
//To fix the error, rename math.pi to math.Pi and try it again.

package main 

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(" My Pi Value is ", math.Pi);
}