package main

import (
	"fmt"
	"time"
)

func switch_conditional() {
	case1, case2 := "B", "A"
	switch val := "A"; val {
	case case1:
		fmt.Println(case1)
	default:
		fmt.Println(case2 + " at default")
	}

	h := time.Now().Hour()
	switch { // equals to true; can replace long if-else chain
	case h < 10:
		fmt.Println("Morning")
	case h < 18:
		fmt.Println("Afternoon/Evening")
	default:
		fmt.Println("Night")
	}
}
