package main

import "os"

func main() {
	switch os.Args[1] {
	case "1":
		pointer_func()
	case "4":
		struct_func()
	}
}
