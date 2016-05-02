package main



import (
	"fmt"
)
// the first two numbers in the
// Fibonacci sequence are either 1 and 1, 
// or 0 and 1, depending on the chosen starting
//  point of the sequence, and each subsequent 
//  number is the sum of the previous two.

func fib(x int) int{
 if x <= 3 {
 	return 1
 }else{
 	return fib(x - 1) + fib(x - 2)
 }
}

func main() {
	fibVar := 6
	x := fib(fibVar)

	fmt.Printf("The %d fibonacci is : %d\n", fibVar, x)
}