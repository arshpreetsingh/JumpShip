package main
import ("fmt")

type person struct {
     name string
     lastname string
     age int
     hobbies []string

}


func main() {

    //hob:=[]string{"ok","bye","thanks"}
    p:=person{name:"arsh",lastname:"singh",age:24,hobbies:[]string{"ok","bye","thanks"}}
    fmt.Println(&p)
}
