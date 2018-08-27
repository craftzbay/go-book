package main

import (
	"fmt"
	"regexp"
)

func main() {
	colorText:="улаан,ногоон шар;цэнхэр"
	re := regexp.MustCompile("[,;\\s]")
	fmt.Printf("%q\n", re.Split(colorText, -1))
}