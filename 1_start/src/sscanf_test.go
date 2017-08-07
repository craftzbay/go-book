package main
import "fmt"

func main() {
 s:="10 20 30"
 a:=0
 b:=0
 c:=0
 fmt.Sscanf(s, "%d", &a, &b, &c)
 println(a, b, c)
}