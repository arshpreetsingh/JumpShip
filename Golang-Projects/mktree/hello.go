package main

import (
	"fmt"
	"time"
)

func run_tree() {
	time.Sleep(30 * time.Second)
	fmt.Println("hello")
}

func main2() {
	var value int
	for {
		go run_tree()
		value += 1
		fmt.Println("this is value", value)
	}

}
