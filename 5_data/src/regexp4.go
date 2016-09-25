package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("i")
	fmt.Println(re.ReplaceAllString("sift rise", "o"))
}