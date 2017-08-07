package main
import "fmt"
var (
 height int   /* гурвалжны өндөр
 width int    /* гурвалжны өргөн */
 area int     /* талбай */
)

func main() {
    fmt.Printf("Өндөр? ")
    fmt.Scanf("%d %d", &width, &height)

    area = (width * height) / 2
    fmt.Printf("Талбай нь %d\n", area)
}
