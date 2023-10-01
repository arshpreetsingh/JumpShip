package main
import ("fmt" 
   "time" )

func routine(value int)(int){
    for i:=0; i<value; i++ {
        fmt.Println(i)
}
    return 10
}


func main(){

    routine(4)
    go routine(10)
    go routine(20)
    time.Sleep(time.Second)
}
