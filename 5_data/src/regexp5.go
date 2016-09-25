package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(\w+),\s*(\w+)`)
	fmt.Println(re.ReplaceAllString("one, two", "$2, $1"))
}