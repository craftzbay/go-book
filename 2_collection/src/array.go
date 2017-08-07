package main
import "fmt"

func main() {
    data:=[]float32{ 34.0, 27.0, 45.0, 82.0, 22.0 }

    total:= data[0] + data[1] + data[2] + data[3] + data[4]
    average:=  total / 5.0
    fmt.Printf("Нийт %f Дундаж %f\n", total, average)
}