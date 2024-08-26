package main

import (
	"fmt"
	"runtime"
)

func main() {
	version := runtime.Version()
	fmt.Printf("Go Version: %s\n", version)
}
