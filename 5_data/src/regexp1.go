package main

import "fmt"
import "regexp"

func extractNumbers(str string) { 
  re,_:=regexp.Compile("[+-]?\\d+\\.?\\d*") 
  all := re.FindAllString(str, -1) 
  fmt.Printf("%q", all) 
} 

func  main() {
	extractNumbers("Өнөөдөр 9 сарын 1")	
}
