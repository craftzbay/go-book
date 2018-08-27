package main
import "fmt"

func main() {
  var c int
  /* жижиг үсгүүдийн код */
  for c = 'a'; c <= 'z'; c++ {
     fmt.Printf( "%c = %d\n", c, c)
  }
     
  /* том үсгүүдийн код */
  for c = 'A'; c <= 'Z'; c++ {
     fmt.Printf( "%c = %d\n", c, c)
  }
}