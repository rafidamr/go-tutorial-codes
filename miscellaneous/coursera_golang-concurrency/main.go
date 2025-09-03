package main

/*
This code below demonstrates the race condition in Go.
This race condition occurs due to the non-deterministic ordering
of the interleaving concurrent tasks, in this the case the print 
and increment functions. The scheduler may prefer to increase the
value of x before printing it. A communication also occur during this
race condition because these two functions share the same variable.
As a result, the value of x may be printed as 1 or 2 in the console.
*/

import (
	"fmt"
	"time"
)

func increment(x *int) {
	*x = *x + 1
}

func print(x *int) {
	fmt.Println(*x)
}

func main () {
	x := 1
	go print(&x)
	go increment(&x)
	time.Sleep(100 * time.Millisecond) // wait before all goroutines finish
}