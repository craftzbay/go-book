package main

func sum(a []int) int { 
  if len(a) == 1 {
    return a[0]
  }
  return a[0] + sum(a[1:])
}

func main() {
  a:=[]int{1,8,3,2}
  println(sum(a))
}