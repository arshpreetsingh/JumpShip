package main
import (
    "fmt"
    "time"
)

func worker(c chan bool, value int){
    fmt.Println("Worker started!",value)
    time.Sleep(time.Second*5)
    fmt.Println("Fnished!!")
    c<-true

}


func main(){
    mychan:=make(chan bool)
    go worker(mychan,10)
    go worker(mychan,20)
    <-mychan
    <-mychan

}
