package main

import "fmt"

type Box struct {
  name	string		/* барааны нэр */
  quantity	int		/* тоо ширхэг */
  cost	float64	/* нэг бүрийн үнэ */
}


func main() {
  b := Box {"Алим", 100, 500.0}
  fmt.Println(b)
}
