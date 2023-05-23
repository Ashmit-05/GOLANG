package main

import "fmt"

func main() {
	fmt.Println("Result : ", adder(3, 5))

	proResult1, proResult2 := proAdder(1, 2, 3, 4, 5, 6)
	fmt.Println("Pro result : ", proResult1)
	fmt.Println("Message from pro result : ", proResult2)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Added all the values"
}
