package main

import "fmt"

func ComputeSum(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	return total
}

func main() {
	fmt.Println("Starting computation")
	res := ComputeSum(10)
	fmt.Printf("Result: %d\n", res)
	if res == 45 {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}
