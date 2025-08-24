package main

import "os"

func main() {
	switch os.Args[1] {
	case "11":
		basic_types()
	case "13":
		type_conversion()
	}
}
