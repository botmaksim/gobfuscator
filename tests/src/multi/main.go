package main

import "fmt"

func main() {
	fmt.Println("Multi-file project started!")
	a := 10
	b := 20
	sum := AddNumbers(a, b)
	PrintResult("Sum", sum)
	
	product := MultiplyNumbers(a, b)
	PrintResult("Product", product)
}
