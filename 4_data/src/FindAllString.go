package main
import (
 "fmt"
 "regexp"
)

func extractNumbers(str string) {
  re,_:=regexp.Compile("[+-]?\\d*\\.?\\d*")
  all := re.FindAllString(str, -1)
  fmt.Println("FindAllString = ", all)
}

func main() {
  extractNumbers("98.0 +12.1 hello 11")
}