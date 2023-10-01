package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sync"
	"time"
)

// The type of data we want ot process.
type item struct {
	price    int
	category string
}

func read_data(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()
	var hashData []string
	if err != nil {
		fmt.Println("Unable to Read File")
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	// fmt.Println("this is data!", reflect.TypeOf(fileScanner))
	for fileScanner.Scan() {
		hashData = append(hashData, fileScanner.Text())
	}
	return hashData
}

func main() {
	hashData := read_data("input.txt")

	done := make(chan bool)
	defer close(done)
	// This function will generate Items!
	// Passing items into chanels.

	///====================================================
	// BAsed on No. of Items we want to Process we will create those many Channels!

	/// Esnu Handle karan lai ikk function banao and dekho how to operate that!!!

	w1 := dataGen(hashData)
	w2 := dataGen(hashData)
	finalData := FanInhash(done, w1, w2)
	fmt.Println("this is final_data", <-finalData)
	fmt.Println("this is final_data", <-finalData)
	fmt.Println("this is final_data", <-finalData)
	fmt.Println("this is final_data", <-finalData)
	fmt.Println("this is final_data", <-finalData)

	c := gen(item{8, "shirt"},
		item{3, "pent"},
		item{2, "socks"},
		item{9, "joggers"},
		item{9, "pent"},
	)

	fmt.Println("this is type of C which is gen() output", reflect.TypeOf(c), len(c))
	// <-chan main.item size->5
	//

	// gen() will create chanel that will receive 10 values.
	// discount() will be function that will calculate the Hash() of each string!
	// fanin() will return data came from 100 channels, which is [100]string

	//c1 := discount(done, c)
	// c2 := discount(done, c)
	// ^^^^^^^Passing two routines above means if one is blocking, other will return Data!!

	// Here in the below Example you need to check/see Fan-in is the one where you
	// have to Run 100s of goroutines!!

	// for _, c := range channels {
	//fmt.Println(reflect.TypeOf(c)) // c is <-chain item
	//fmt.Println("this is C", c)
	//	go output(c)
	//}

	// Then run same process again by combining two items.
	// Something like that.....
	// i,i+2

	// then pass it again to gen() function, Until
	// len(data)==1

	//
	// some notes from Reflect
	//reflect.ValueOf(x interface{}).
	//reflect.TypeOf(x interface{}).
	//Type.Kind().
	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.ValueOf(c))
	// kind := reflect.TypeOf(c)
	// fmt.Println(kind.Kind())
	// for i := range c {
	// 	i := reflect.TypeOf(i)
	// 	fmt.Println(i.Kind())
	// }

	// Passed chanel generator into the function to be processed
	//out := discount(c) // till here discount() is just sittng idle!
	//fmt.Println(out)
	//for processed := range out {
	//	fmt.Println("This is the value cming out from chanel", processed)
	//}
	// Now create a Function which will take all chanels as input!

	//

	c1 := discount(done, c)
	c2 := discount(done, c)
	// for i := range c1 {
	// 	fmt.Println("coming from c1", i)
	// }
	// for i := range c2 {
	// 	fmt.Println("coming from c2", i)
	// }
	//fmt.Println("this will be C1 and C2", reflect.TypeOf(c1), reflect.TypeOf(c1))
	// Now we have Multiple Channels,
	// One channel Contain Multiple Items!
	data := FanIn(done, c1, c2)
	fmt.Println(<-data)
	//data is one chanel containing multiple items! (<-chan main.item)
	//fmt.Println("this will be return of Fanin() function", reflect.TypeOf(data))
	// for i := range data {
	// 	//fmt.Println("This is Type of Fanout-data")
	// 	//fmt.Println(reflect.TypeOf(i)) //// this will return TypeOf-item!!
	// 	//kind := reflect.TypeOf(i)
	fmt.Println(<-data)
	// 	//fmt.Println("This is Kind of C", kind.Kind())
	// 	fmt.Println(i)
	// }

}

// Number of Channels you want to Open-Up in the system.
// Let's suppose one Chanel is Blocking, other chanel will stil lb operational.

func FanInhash(done <-chan bool, channels ...<-chan string) <-chan string {

	var wg sync.WaitGroup
	out := make(chan string)
	output := func(c <-chan string) {
		defer wg.Done()
		// Iterate Each Item from c, send back to (item-Chan)
		for i := range c {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		// It will Run Twicw Now!!
		//fmt.Println(reflect.TypeOf(c)) // c is <-chain item
		//fmt.Println("this is C", c)
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func FanIn(done <-chan bool, channels ...<-chan item) <-chan item {
	var wg sync.WaitGroup
	out := make(chan item)

	output := func(c <-chan item) {
		defer wg.Done()
		// Iterate Each Item from c, send back to (item-Chan)
		for i := range c {
			fmt.Println("this is i from range", i)
			select {
			case out <- i:
				fmt.Println("this is out->", i)
			case <-done:
				fmt.Println("this is done->", done)
				return
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		// It will Run Twicw Now!!
		//fmt.Println(reflect.TypeOf(c)) // c is <-chain item
		//fmt.Println("this is C", c)
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Here we are taking Just one Chanel as input
// one chanel of type item make(chan item,len(items))
// Buffered Chanels!
func GetHash(done <-chan bool, hashData <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := range hashData {
			//fmt.Println(i.category)
			i := i + "this is my hash!!"
			fmt.Println("this is i now!!", i)
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}

func discount(done <-chan bool, items <-chan item) <-chan item {
	out := make(chan item)
	go func() {
		defer close(out)
		for i := range items {
			//fmt.Println(i.category)
			if i.category == "pent" {
				fmt.Println("Going for Pent!!")
				time.Sleep(1 * time.Second)
				i.price /= 2
			}
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}

func dataGen(hashData []string) <-chan string {
	out := make(chan string, len(hashData))
	for _, i := range hashData {
		out <- i
	}
	close(out)
	return out

}

func gen(items ...item) <-chan item {

	// declared ChaneL of type ITEMS
	out := make(chan item, len(items))
	//fmt.Println("this is Chanel", out)
	// passing values to Chan
	for _, i := range items {
		out <- i
	}
	close(out)
	return out
}
