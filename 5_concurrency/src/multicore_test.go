package main

import (
 "time"
)


func doJob(quit chan int) {
    // тодорхой ажил гүйцэтгэхийг оруулав
    time.Sleep(time.Millisecond * 200)
    quit <- 1
}

func main() {
    routineQuit := make(chan int)
    
    for i:=0; i<50; i++ {
      go routine(routineQuit)
    }

    // бүх функцээс дууссан дохио хүлээх
    for i:=0; i<50; i++ {
     <-routineQuit 
    }
}