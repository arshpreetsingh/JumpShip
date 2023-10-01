package main
import (
    "fmt"
)

func myfunc(a int, b string) (int,string){
    return a,b 
}

func main(){
    a,b:=myfunc(1,"hello")
    fmt.Println(a,b)


}
