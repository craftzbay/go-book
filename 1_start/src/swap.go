package main 
import "fmt" 

var a int
var b int 

func main() { 
  fmt.Printf("a b ? ") 
  fmt.Scanf("%d %d", &a, &b) 

  fmt.Printf( "a=%d, b=%d\n", a, b ) 
  a^=b; b^=a; a^=b 
  fmt.Printf( "a=%d, b=%d\n", a, b ) 
}