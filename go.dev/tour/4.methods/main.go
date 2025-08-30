package main

import "os"

func main() {
	switch os.Args[1] {
	case "1":
		methods_func()
	case "9":
		interfaces_func()
	case "14":
		interfaces2_func()
	case "18":
		stringer_func()
	case "20":
		error_func()
	case "22":
		reader_func()
	}
}
