package main

import (
	asciiart "asciiart/package"
	"os"
)

func main() {
	asciiart.Program(os.Args[1:])
}
