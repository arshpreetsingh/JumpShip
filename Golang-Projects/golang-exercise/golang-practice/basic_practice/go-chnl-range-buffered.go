package main

import ("fmt")



func main(){
    mychan:=make(chan string,3)
    
    mychan<-"one"
    mychan<-"two"
    mychan<-"three"
    close(mychan)
    for value:=range mychan{
        fmt.Println(value)

}

}
