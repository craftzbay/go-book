package main
import (
 "fmt"
 "regexp"
)

func IsIP(str string) bool {
  m,_:=regexp.MatchString("^([0-9]{1,3}\\.){3}[0-9]{1,3}$", str)
  return m
}

func main() {
  fmt.Println(IsIP("192.168.1.100"))
}