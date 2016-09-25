package main

import (
	"fmt"
	"regexp"
)

func main() {
	txt := "Та ubs121@gmail.com, ub121@hotmail.com хаягуудаар холбоо барьж болно."

	re, _ := regexp.Compile(`\w+@\w+\.\w+`)
	all := re.FindAllString(txt, -1)

	for _, m := range all {
		fmt.Printf("%s\n", m)
	}
}
