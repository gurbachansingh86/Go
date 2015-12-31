// Switch
// You probably knew what switch was going to look like.
//
// A case body breaks automatically, unless it ends with a fallthrough statement.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
