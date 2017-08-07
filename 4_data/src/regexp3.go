package main

import "regexp"
import "fmt"

func main() {
 	txt :="Даваа Дээд:26 Доод:21 Мягмар Дээд:23 Доод:20"


	re,_:=regexp.Compile(`([а-яА-Я]+)\s*Дээд:(\d+)\s*Доод:(\d+)`) 
	all := re.FindAllStringSubmatch(txt, -1) 
  	
  	for _, m := range all {
		fmt.Printf("%s: %s-%s\n", m[1], m[2], m[3]) 
  	}
  	
}
