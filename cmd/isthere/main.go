package main

import (
	"fmt"
	"os"

	"github.com/tuxikus/isthere"
)

func help() {
	fmt.Println("usage: isthere <command>")
	fmt.Println()
	fmt.Println("options:")
	fmt.Println("-h, --help:    print the help")
	fmt.Println("-v, --version: print the version")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Illegal argument count")
		help()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-h", "--help":
		help()
		os.Exit(0)
	case "-v", "--version":
		fmt.Println(isthere.GetVersion())
		os.Exit(0)
	}

	command := os.Args[1]
	isthereOutput, err := isthere.IsThere(command)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(isthereOutput)
	os.Exit(0)
}
