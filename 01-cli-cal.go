package main

import (
	"fmt"
	"os"
)

func main() {
	// First part of os.Args is always the binary exe file path
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Arg number %d is %v\n", i, os.Args[i])
	}
}
