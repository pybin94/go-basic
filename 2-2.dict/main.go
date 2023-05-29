package main

import (
	"dict/mydict"
	"fmt"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	// Didn't check duplicate
	// dictionary["hello"] = "hello"

	baseWord := "hello"
	err := dictionary.Add(baseWord, "First")
	if err != nil {
		fmt.Println(err)
	}

	dictionary.Update(baseWord, "Second")

	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	dictionary.Search(baseWord)

	definition, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
