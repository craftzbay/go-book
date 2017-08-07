package main
import "fmt"

var (
  total int      /* нийт дүн */
  current int    /* хэрэглэгчээс өгсөн утга */
  counter int    /* давталтын тоолуур */
)

func main() {
    total = 0
    
    for counter = 0; counter < 5; counter++ {
        print("Тоо? ")
        fmt.Scanf("%d", &current)
        total += current
    }
    
    fmt.Printf("Нийт дүн %d\n", total)
}