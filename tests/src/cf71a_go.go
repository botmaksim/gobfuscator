package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		if len(s) > 10 {
			fmt.Printf("%c%d%c\n", s[0], len(s)-2, s[len(s)-1])
		} else {
			fmt.Println(s)
		}
	}
}
