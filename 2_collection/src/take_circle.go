package main

func TakeCircle(x int, y int, step int) bool {
   /* Аль нэг шонгийн цагираг дууссан бол 
	хожил/хожигдолыг тооцоход хялбархан */	
   if x == 0 || y == 0 {
     return ((x + y) % 2 == step % 2)
   } else { /* эсрэг тохиолдолд рекурсивээр цагираг авах үйлдлийг давтана */
	if step % 2 == 0 {
	   // A тоглогч цагираг авах хувилбарууд
	   return (TakeCircle(x - 1, y, step + 1) &&
		TakeCircle(x - 1, y - 1, step + 1)&& 
		TakeCircle(x, y - 1, step + 1))
	} else {
	   // B тоглогч цагираг авах хувилбарууд
	   return (TakeCircle(x - 1, y, step + 1) ||
		TakeCircle(x - 1, y - 1, step + 1)|| 
		TakeCircle(x, y - 1, step + 1))
	}
   }
}

/* Функцийг турших үндсэн програм */
func main() {
   /* (7,7) хосын хувьд А тоглогч хожих эсэхийг шалгах */
   println( "A тоглогч хожих уу? ") 
   if TakeCircle(7,7,1) {
     println("тийм")
   } else {
     println("үгүй") 
   }
}