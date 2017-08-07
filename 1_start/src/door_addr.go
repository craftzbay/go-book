package main
import "fmt"

var (
  floors uint 	/* давхарын тоо */
  flatPerFloor uint /* нэг давхарт суудаг айлын тоо */
  n uint 	/* хайх хаалганы дугаар */
)

func main() { 
  fmt.Printf("Байрны давхарын тоо? ")
  fmt.Scanf("%d", &floors)

  fmt.Printf("Нэг давхарт суудаг айлын тоо? ")
  fmt.Scanf("%d", &flatPerFloor)
  
  fmt.Printf("Хайх хаалга? ")
  fmt.Scanf("%d", &n)
  
  /* Нэг орцон дахь айлын тоог тооцоолох */
  flatPerEntrance := floors * flatPerFloor

  /* Орцны дугаарыг олох */
  entrance := n / flatPerEntrance;
  if n % flatPerEntrance > 0 {
   entrance++
  }

  /* Давхарын дугаарыг олох */
  temp := n -(entrance - 1) * flatPerEntrance
  fl := temp / flatPerFloor
  if temp % flatPerFloor > 0 {
   fl++
  }

  fmt.Printf( "Орц = %d\n", entrance)
  fmt.Printf( "Давхар = %d\n", fl )
}