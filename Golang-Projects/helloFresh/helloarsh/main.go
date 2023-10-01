/**
Common description:
You need to write a program that gets webpages, extracts content
length of a page and puts this value into the map.
When all address processed, program prints all pairs (webpage address, content length).

Specific requirements:

Worker is a function that receives webpage address (ex: "google.com") and gets page content.
Worker gets only basic html document and do not receive external resources like images, css and so on.
You need to create two workers (two go routines).
You need to use channel(s) to send webpage address to the workers.
On success, if it's possible to get webpage, put the length of
webpage content into rusults.ContentLength map.
Use as a key webpage address, and content length as the value.
On failure, if it's impossible to get webpage because some errors, put into the results.
ContentLength webpage address as a key, and -1 as a value.

When all webpages from webPages slice processed, print each key-value from webPages.
Example:
google.com - 4501
...


If the program runs more than 30 seconds, you must stop all workers
(go routines), print message "out of time" and exit.


You can modify provided sources as you wish. You can extend the results structure as you need
just keep two provided fields in it.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	webPages = []string{
		"yahoo.com",
		"google.com",
		"bing.com",
		"amazon.com",
		"github.com",
		"gitlab.com",
	}

	results struct {
		// put here content length of each page
		ContentLength map[string]int

		// accumulate here the content length of all pages
		TotalContentLength int
	}
)

func httpWorker(urlList []string, data chan map[string]int) {
	results.ContentLength = make(map[string]int)
	for _, url := range urlList {
		resp, err := http.Get("https://" + url)
		if err != nil {
			fmt.Println(err)
		}
		if resp.StatusCode == 200 {
			var content []byte
			content, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			results.ContentLength[url] = len(content)
		} else {
			results.ContentLength[url] = -1
		}
	}
	data <- results.ContentLength
}

func main() {
	datalist1 := make(chan map[string]int)
	datalist2 := make(chan map[string]int)
	defer close(datalist1)
	defer close(datalist2)
	webPagesList1 := webPages[0 : len(webPages)/2]
	webPagesList2 := webPages[len(webPages)/2:]
	go httpWorker(webPagesList1, datalist1)
	go httpWorker(webPagesList2, datalist2)
	select {
	case results.ContentLength = <-datalist1:
		totalCount := 0
		for _, value := range results.ContentLength {
			totalCount += value
		}
		results.TotalContentLength = totalCount
		fmt.Println(results)
	case <-time.After(30 * time.Second):
		fmt.Println("Can't Wait that Long Now, Gotta GO!")
	}
}
