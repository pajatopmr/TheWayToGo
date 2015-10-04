// Create a non-recursive version of the Fibonacci implementation.
package main

import "fmt"

func fib(n int) {
	switch {
	case n == 0:
		fmt.Println("0:1")
	case n == 1:
		fmt.Println("0:1 1:1")
	default:
		fmt.Print("0:1 1:1")
		n1 := 1
		n2 := 1
		for i := 2; i < n; i++ {
			val := n1 + n2
			fmt.Printf(" %d:%d", i, val)
			n1 = n2
			n2 = val
		}
		fmt.Println()
	}
}
