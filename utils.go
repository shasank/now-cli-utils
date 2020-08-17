package utils

import (
	"fmt"
)

// Constants
const (
	ver = "v1.0"
)

func main() {
	fmt.Println(GetVersion())
}

// GetVersion ->
func GetVersion() string {
	return ver
}
