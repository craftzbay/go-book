package main

import "fmt"

var year int 	/* он */
var month int 	/* сар */

func main() { 
 fmt.Print("Он, сар ? ")
 _, err := fmt.Scanf("%d %d", &year, &month)
 
 if err != nil {
   fmt.Println("Он сар алдаатай байна!")
   return
 }
 
 fmt.Print(year, month, " сар ")

 switch month {
  case 1,3,5,7,8,10,12: 
	fmt.Println("31 хоногтой")
  case 4,6,9,11: 
	fmt.Println( "30 хоногтой" )
  case 2: // өндөр жил эсэх
	if year == 400 || (year % 4 == 0 && year % 100 == 0) {
  	   fmt.Println( "29 хоногтой" )
	} else {
	   fmt.Println( "28 хоногтой" )
        }
 }
}