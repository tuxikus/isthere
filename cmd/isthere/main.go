package main

import (
	"fmt"
	"os"

	"github.com/tuxikus/isthere"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Illegal argument count")
		fmt.Println("usage: isthere <command>")
		os.Exit(1)
	}

	command := os.Args[1]
	fmt.Println(isthere.IsThere(command))
	os.Exit(0)
}
