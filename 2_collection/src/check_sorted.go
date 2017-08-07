package main

import "fmt"

var n int 	/* массивын элементийн тоо */

func main() {
  fmt.Printf("Элементийн тоог оруулна уу n? ")
  fmt.Scanf("%d", &n )

  /* n-ээс хамаарсан хувьсах урттай массив зарлах */
  array:=make([]int, n)

  /* Элементүүдийг нэг нэгээр нь оруулах */
  for i:=0; i<n; i++ {
    fmt.Printf("array[%d]=", i)
    fmt.Scanf("%d", &array[i] )
  }

  /* Массивыг хэвлэж харуулах */
  fmt.Printf("\narray[")
  for i:=0; i<n-1; i++ {
    fmt.Printf( "%d, ", array[i] )
  }
  fmt.Printf( "%d] ", array[n - 1] )

  /* Элементүүд буурах эрэмбээр байрласан эсэхийг шалгах */
  ind := 1
  for ind < n && array[ind - 1] > array[ind] {
   ind++
  }

  if ind >= n {
   println("<= буурах эрэмбэтэй массив мөн")
  } else {
   println("<= буурах эрэмбэтэй массив биш байна!")
  }
}