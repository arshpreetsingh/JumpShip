package main

import ("fmt")

type payload struct{
    url string
    calltype string 

}



func (p payload) httpcall() *payload {
    //fmt.Println(p)
    return &p 
}

func main(){
    fmt.Println("aka oops in golang")
    p:=payload{url:"http:/google.com",calltype:"post"}
    fmt.Println(p.httpcall())


}
