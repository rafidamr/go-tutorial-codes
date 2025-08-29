package main

import "os"

func main() {
	switch os.Args[1] {
	case "1":
		looping()
	case "8":
		newtonmethod_func()
	case "9":
		switch_conditional()
	case "12":
		defer_func()
	}
}
