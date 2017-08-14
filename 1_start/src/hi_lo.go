package main

import "fmt"

func main() {
	var n byte
	fmt.Printf("Бүхэл тоо оруулна уу? ")
	fmt.Scanf("%d", &n)

	fmt.Printf("Дээд 4 бит = %d\n", (n&0xF0)>>4)
	fmt.Printf("Доод 4 бит = %d\n", (n & 0x0F))
}
