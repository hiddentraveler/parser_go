package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./examples/oo.lang")
	source := string(bytes)

	fmt.Printf("code: %s\n", source)
}
