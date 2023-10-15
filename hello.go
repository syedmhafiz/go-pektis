package main

import (
	"fmt"
)

func main() {

	if i := 1; i < 3 {
		fmt.Println(i * 10)
	} else {
		i += 10
		fmt.Println(i)
	}
}
