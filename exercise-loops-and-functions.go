package main

import (
	"fmt"
)

/** First Approach ***/
func Sqrt(x float64) float64 {
	return x*x;	
}

/** Second Approach 
func Sqrt(x int)(s int) {
	s = 1
	for i := 1; i <= 2; i++ {
		s = s * x		
	}	
	return s;
}**/

func main() {
	fmt.Println(Sqrt(0))
}
