package main

import (
	"fmt"
)

func main () {
	cetakAngka(1,2,3,5,6,6,7,9)
}

func cetakAngka (inp ...int) int {
	if len(inp) == 0 {
		return 0
	}
	for _, x := range inp {
		fmt.Printf("%d\n", x)
	}
	return 0
}