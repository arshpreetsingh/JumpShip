package main
import (
        "fmt" 
        "time"
   )

//id int, jobs <-chan int, results chan<- int)
func myWorker(job <-chan int, result chan<- int, jobid int){

    time.Sleep(time.Second*3)
    fmt.Println("this is job-ID",jobid)
    time.Sleep(time.Second*1)
    fmt.Println("this is JOB", job)
    time.Sleep(time.Second*3)
    fmt.Println("sending result back!!")
    result<-28

}

func main(){

    job:=make(chan int,10)
    result:=make(chan int,10)

    // now start required number of Wokers!
    for i:=0;i<3;i++{
        go myWorker(job,result,i)
}


    // Now we have to send values to job's Chanel(receiver Chanel)

     for i:=0;i<10;i++{
          job<-i
}
    close(job) 

    // Now we have to receive values from Chanel//

    for i:=0;i<10;i++{
        <-result
}

}
