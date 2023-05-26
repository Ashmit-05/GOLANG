package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer deferExample()
	fmt.Println("Hello")
}

func deferExample() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
