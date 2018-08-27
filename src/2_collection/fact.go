package main

func fact(number int) int { 
  if number == 0 { 
    return (1)
  }
  return number * fact(number-1) 
}

func main() {
  println(fact(3))
}