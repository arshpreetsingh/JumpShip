package main

import (
    "fmt"
)

func main(){
    mymap:=make(map[string]int)
    mymap["arsh"]=1
    mymap["pawan"]=2
    mymap["amber"]=3
    fmt.Println(mymap)
    secondmap := map[string]int{"arsh":1,"amber":2}
    fmt.Println(secondmap)
}
