package main
import (
    "fmt"
)

func main() {
    s:=make([]int,10) 
    for i:=0;i<10;i++{
         s[i]=i

    }
    fmt.Println(s)
}
