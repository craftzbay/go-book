package main
import "fmt"

func main() {
    data:=make([]int, 5)
    seven_count:=0 // 7-н тоо
    three_count:=0 // 3-н тоо
    
    fmt.Printf("5 ширхэг тоо оруул\n")
    fmt.Scanf("%d %d %d %d %d", 
      &data[0], &data[1], &data[2], &data[3], &data[4])

    for index:=1; index<5; index++ {
        if data[index] == 3 {
            three_count++
        }

        if data[index] == 7 {
            seven_count++
        }
    }
    fmt.Printf("Гуравууд %d Долоонууд %d\n", three_count, seven_count)
}