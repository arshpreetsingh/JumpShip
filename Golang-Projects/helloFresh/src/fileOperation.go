// Package src /**/
package src

import (
	"encoding/json"
	"log"
	"os"
)

// ReadJson Will Read json-input Data and Return iterator.
func ReadJson(fileName string) *json.Decoder {
	log.Println("Reading Json File:", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Not Able to Read File: Error:", err)
	}
	decoder := json.NewDecoder(file)
	decoder.Token()
	log.Println("Reading JSON File Completed!")
	return decoder
}
