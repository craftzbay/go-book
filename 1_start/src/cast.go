package main

import "fmt"

var won, lost int = 5, 3
var anything interface{} = 123

func main() {

	won = 5
	lost = 3
	//ratio = won / lost /* харьцаа нь 1.0 (буруу үр дүн) */
	/* Дараах тохиолдолд харьцааг зөв тооцоолно */
	ratio := float32(won) / float32(lost)

	fmt.Printf("Харьцаа=%f\n", ratio)

	switch v := anything.(type) {
	case string:
		fmt.Println(v)
	case int, int32, int64:
		fmt.Println(v)
	default:
		fmt.Println("unknown")
	}
}
