package main
import "fmt"

func main() {
    var value int
    fmt.Print("Утга оруул: ")
    fmt.Scanf("%d", &value)

    fmt.Printf("%d * 2 = %d\n", value, value * 2)
}