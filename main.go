package main

import (
	"fmt"
)

// Constants
const (
	ver = "v1.0"
)

func main() {
	fmt.Println(getVersion())
}

func getVersion() string {
	return ver
}
