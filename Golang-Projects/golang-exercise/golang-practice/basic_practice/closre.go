package main

import (
    "fmt"
)

func outerfunc(mesage string)  func(string) {
     return func(message string) string {
      return("hello"+message)
      }
}


func main() {
    output:=outerfunc("arshpreet singh khangura")
    fmt.Println(output)
}
