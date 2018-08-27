package main

import "fmt"

var (
	month     int
	day       int
	totalDays int /* нийт өдрийн тоо */
	monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

/* АНХААР:  Энэ програмд өндөр жилийг тооцоогүй. 2-р сарыг 28 хоногтой гэж авав  */
func main() {
	fmt.Printf("Сар/Өдөр ? ")
	fmt.Scanf("%d/%d", &month, &day)

	totalDays = day

	for i := 0; i < month-1; i++ {
		totalDays += monthDays[i]
	}

	fmt.Printf("Жилийн эхнээс өнгөрсөн өдрийн тоо: %d\n", totalDays)
}
