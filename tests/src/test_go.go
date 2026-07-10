package main

import (
	"fmt"
)

func main() {
	secretKey := "SuperSecretKey123"
	myVariable := 42

	fmt.Println("My logic runs:", secretKey)
	
	if myVariable == 42 {
		fmt.Println("Logic is intact!")
	}
}
