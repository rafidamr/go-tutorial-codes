package main

import "os"

func main() {
	switch os.Args[1] {
	case "1":
		firstFunc()
	case "2":
		secondFunc()
	}
}
