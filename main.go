package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pathVal := os.Getenv("PATH")
	for _, str := range strings.Split(pathVal, ":") {
		fmt.Println(str)
	}
}
