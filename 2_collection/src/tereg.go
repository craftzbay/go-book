package main

import "fmt"

func main() {
	board := [][]int{{0, 1, 1, 1}, {1, 0, 4, 3}, {0, 1, 3, 5}, {0, 0, 2, 5}} // оролтын өгөгдөл
	n := len(board[0])                                                       // хөлгийн хэмжээ

	occupiedRows := make(map[int]bool) // тэрэг байрласан мөрүүд
	occupiedCols := make(map[int]bool) // тэрэг байрласан баганууд

	// r,c байрлалын нийлбэрийг олох функц
	posSum := func(r, c int) int {
		sum := 0
		for x := 0; x < n; x++ {
			if !occupiedCols[x] { // өмнө байрласан баганыг тооцохгүй
				sum += board[r][x]
			}
			if !occupiedRows[x] { // өмнө байрласан мөрийг тооцохгүй
				sum += board[x][c]
			}
		}
		sum = sum - 2*board[r][c] // r,c байрлал дээрх утгыг хасах
		return sum
	}

	totalSum := 0

	for x := 0; x < 2; x++ { // 2 ширхэг тэрэг байрлуулна
		max := 0
		r := 0
		c := 0

		// max утгатай байрлал олох
		for i := 0; i < n; i++ {
			if !occupiedRows[i] {
				for j := 0; j < n; j++ {
					if !occupiedCols[j] {
						tmpSum := posSum(i, j)
						if max < tmpSum {
							max = tmpSum
							r = i
							c = j
						}
					}
				}
			}
		}

		totalSum += max
		occupiedRows[r] = true
		occupiedCols[c] = true

		fmt.Printf("Pos (%d, %d), Sum %d\n", r, c, max)
	}

	fmt.Printf("Total sum %d\n", totalSum)
}
