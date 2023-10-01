package main

import (
    "fmt"
     "encoding/csv"
     "os"
      "flag"
      "strings"
      "time"
)


type problem struct {
    question string
    answer   string
}

func ParseLines(lines [][]string) []problem {
 //// This function will take lines those we will read from files
//// and turn those into a problem{} 
     ret:=make([]problem,len(lines)) // we need to declare size of []problem array
     for i,line := range lines {
     ret[i]=problem {
            question:strings.TrimSpace(line[0]),
            answer:strings.TrimSpace(line[1]),
     }
    }
    return ret
}

func main() {
    csvFilename := flag.String("csv", "problem.csv file", "File that you want to read")
    timeLimit := flag.Int("time_limit",30,"enter Your Time Limit")
    //fmt.Println(timelimit)
    flag.Parse()

    file,err := os.Open(*csvFilename)
    if err!=nil {
         fmt.Println("unable to read file")
         os.Exit(1)
                 }
     data :=csv.NewReader(file)
     lines,err:=data.ReadAll()
     if err!=nil{
         fmt.Println(err, "eror while eading CSV files")
     }
    problem:=ParseLines(lines)
    timer:=time.NewTimer(time.Duration(*timeLimit)*time.Second)
    counter := 0
    //ToDo: 1. Print Question in Standard Output,2. Accept answer from
    // user and check, 3. have counter of each question, 4. Add strings.TrimSpaces()
    for _,value:=range problem {
        fmt.Println("Please answer", value.question)
        answerChanel:=make(chan string)
        go func() {
        var my_answer string
        fmt.Scanf("%s", &my_answer)
        answerChanel<-my_answer
        }()

        select {
        case <-timer.C:
        fmt.Println("Your time is Expired!")
        fmt.Println("Your Score is",counter,"out of", len(problem))
        return
        case my_answer:=<-answerChanel:
        if my_answer==value.answer {
            fmt.Println("Correcet!")
            counter++ }

}
}

}
