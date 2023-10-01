package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"reflect"

	"github.com/cbergoon/merkletree"
)

//TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	x string
}

//CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (t TestContent) Equals(other merkletree.Content) (bool, error) {
	return t.x == other.(TestContent).x, nil
}

func main() {
	//Build list of Content to build tree
	var list []merkletree.Content
	list = append(list, TestContent{x: "Hello"})
	list = append(list, TestContent{x: "Hi"})
	list = append(list, TestContent{x: "Hey"})
	list = append(list, TestContent{x: "Hola"})

	//Create a new Merkle Tree from the list of Content
	fmt.Println("this is our List", list)
	t, err := merkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
	}

	/*
			   [95 48 204 128 19 59 147 148 21 110 36 178 51 240 196 190 50 178 78 68 187 51
			   129 240 44 123 165 38 25 208 254 188]

		     first half:
		     [209 35 249 125 162 93 164 176 139 150 46 100 184 4 43 218 115 153 80 68 63
		     23 33 50 229 165 12 222 16 231 182 188]
	*/
	//Get the Merkle Root of the tree
	mr := t.MerkleRoot()
	log.Println("this is merketl-Root-2", reflect.TypeOf(mr))

	//Verify the entire tree (hashes for each node) is valid
	vt, err := t.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree
	vc, err := t.VerifyContent(list[0])
	if err != nil {
		log.Fatal(err) // vt, err := t.VerifyTree()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println("Verify Tree: ", vt)
	}

	log.Println("Verify Content: ", vc)

	//String representation
	log.Println(t)
}
