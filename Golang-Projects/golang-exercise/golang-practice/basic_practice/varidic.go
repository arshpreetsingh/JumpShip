package main

import (
    "fmt"

)

func hello(values ...string){
    fmt.Println(values)
}

func main(){
    hello("arsh","amber","ok","bye")

}
