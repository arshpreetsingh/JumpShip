package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		number := rand.Int()
		time.Sleep(time.Second * 2)
		fmt.Println("number", number)

	}
}
