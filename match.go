package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	src := strings.Fields(os.Args[1])
	fmt.Println(len(src))
}
