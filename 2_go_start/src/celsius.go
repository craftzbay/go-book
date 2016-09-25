package main
import "fmt"
/*
 * Энэ програм нь 0-ээс 100 хүртэлх Celsius хэмийг 
 * Fahrenheit хэм рүү хөрвүүлнэ
 */
func main() {
    for celsius:=0; celsius<=100; celsius++ {
        fmt.Printf("Celsius:%d Fahrenheit:%d\n",
            celsius, (celsius * 9) / 5 + 32)
    }
}