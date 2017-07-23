# Жишээ дасгал

Жишээ 1. Өгөгдсөн бүхэл тоон А[N] массив буурах эрэмбээр эрэмбэлэгдсэн эсэхийг шалгах програм бичээрэй.

```
package main
import "fmt"

var n int 	/* массивын элементийн тоо */

func main() {
  fmt.Printf("Элементийн  тоог оруулна уу n? ")
  fmt.Scanf("%d", &n )

  /* n-ээс хамаарсан хувьсах урттай массив үүсгэх */
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
```

Жишээ 2. Өгөгдсөн тоон дарааллыг өсөх дарааллаар эрэмбэлэн харуулах програм бич. (тоонууд сул зайгаар тусгаарлагдан нэг мөрөнд байрлаж өгөгдөнө)

```golang
package main
import ( "bufio" ; "fmt" ; "os" )

var array [200]int

/******************************************************    
 * QuickSort -  Массивын өгөгдсөн мужийг эрэмбэлэх    *
 * 							        *
 * Параметрүүд                                        *
 *   low –- эрэмбэлэх мужийн доод хязгаар.            *
 *   high –- эрэмбэлэх мужийн дээд хязгаар.           *
 *                                                    *
 ****************************************************** /
func QuickSort(low int, high int) {
    i:= low; j:= high
    pivot:=array[(low + high)/2]
    var temp int

    for i<=j {
        for array[i] < pivot { i++ }
        for array[j] > pivot { j-- }
        if i <= j {
            temp = array[i]
            array[i] = array[j]
            array[j] = temp
            i++; j--
        }
    }

    if low < j { QuickSort(low, j) }	/* зүүн хэсгийг эрэмбэлэх */
    if i < high { QuickSort(i, high) } /* баруун хэсгийг эрэмбэлэх */
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    s, _ := reader.ReadString('\n')

    /* элементийн тоо */
    n:=0

    /* Сул зайгаар тусгаарлагдсан  тоон дарааллыг унших */
    for i:=0; i<len(s); n++{
        for i<len(s) && s[i]==' ' { i++ }
        fmt.Sscanf(s[i:], "%d", &array[n])
        for i<len(s) && s[i]!=' ' { i++ }
    }

    /* массивыг эрэмбэлэх */
    QuickSort(0, n - 1)

    /* эрэмбэлэгдсэн  массивыг хэвлэж харуулах */
    for i:= 0; i<n; i++ {
       fmt.Printf("%d ", array[i])
    }
}
```

Жишээ 3. N тооноос К-аар зохиосон бүх боломжит хэсэглэлийг харуулах програм бичээрэй.

```
package main
import "fmt"

var (
  n int		/* нийт элементийн тоо */
  k int		/* хэсэглэж авах элементийн тоо */
)

func main() {
  fmt.Printf("n k ? ")
  fmt.Scanf("%d %d", &n, &k )

  /* хэсэглэлд зориулсан массив */
  a:=make([]int, k+1)

  p:=k
  /* эхний байрлал */
  for i:=1; i<=k; i++ { a[i] = i }

  for p > 0 {
    /* шинэ байрлал */
    for i:=1; i<=k; i++ { print(a[i], " ") }
    println()

    /* дараагийн хэсэглэлийг зохиох */  
    if a[k] == n {
	p--
    } else {
	p = k
    }

    if p > 0 {
      for i:=k; i>=p; i-- {
         a[i] = a[p] + (i - p + 1)
      }
    }
  }
}
```

Жишээ 4.  N ширхэг тоогоор зохиох бүх боломжит сэлгэмлийг харуулах програм бичээрэй.

```
package main
import "fmt"

var (
  a [100]int	/* сэлгэх утгуудыг агуулах массив */
  N int		/* сэлгэх элементийн тоо */
)
func main() {
    print("N=")
    fmt.Scanf("%d", &N)

    /* эхний хувилбар */
    for i:=0; i<N; i++ { a[i] = i + 1 }
    display()

    p:=make([]int, N) /* сэлгэмлийг удирдах массив */
    var (
     j int
     temp int
    )   

    for i:= 1; i < N;  {
        if p[i] < i {
            if i%2>0 {  j = p[i]  } else { j = 0 }
            temp = a[j]; a[j] = a[i]; a[i] = temp;

	    /* сэлгэмлийн  шинэ хувилбарыг хэвлэх */
            display()

            p[i]++
            i = 1
        } else {
            p[i] = 0
            i++
        }
    }
}

// туслах функц
func display() {
    for i:=0; i<N; i++ {  print(a[i], " ") }
    println()
}
```

Жишээ 5. Шатрын мориор нүд бүхэн дээр нэг удаа бууж хөлгийг бүтэн тойрох арга байдаг. Энэ аргыг олох програм зохио.

Шийдэл:  Хөлгийн нүднүүдийг морины нүүдлээр 1,2,3… гэх мэт дугаарлан тэмдэглэж явая. Дугаарлах явцад нүүдэлгүй болж гацаанд орвол нүүдлүүдийг буцаан ухрааж дараагийн боломжит нүүдлийг сонгоно. Энэ маягаар явсаар дугаарлалт 64 хүртэл явж чадвал бүтэн тойрох нэг хувилбарыг оллоо гэсэн үг.

```
package main

var (
 board [8][8]byte
 done bool
)

func main() {
  move(0, 0, 1)
}

func move(x, y int, n byte) {
  if !done && 0<=x && x<8 && 0<=y && y<8 && board[x][y] == 0 {
	board[x][y] = n
	if n == 64 {
	  printBoard()
	  done = true
	  return
	}
	move(x-1,y+2,n+1)
	move(x-2,y+1,n+1)
	move(x-2,y-1,n+1)
	move(x-1,y-2,n+1)

	move(x+1,y-2,n+1)
	move(x+2,y-1,n+1)
	move(x+2,y+1,n+1)
	move(x+1,y+2,n+1)

	board[x][y] = 0
  }
}

func printBoard() {
  for i:=0; i<8; i++ {
    for j:=0; j<8; j++ {
      print(" ", board[i][j])
    }
    println()
  }
}
```
