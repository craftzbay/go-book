package main
import "fmt"

var total int  /* нийт дүн */
var item int   /* нэмэх утга */

func main() {
    total= 0

    for {
        fmt.Printf("Тоо оруулна уу ( 0 оруулвал дуусна) ? ")

        fmt.Scanf("%d", &item)

        if item == 0 {
            break
        }

        total += item
    }

    fmt.Printf("Эцсийн дүн %d\n", total)    
}
