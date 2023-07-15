package main

import (
	"fmt"
	"nmc/dictionary/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "First"}
	fmt.Println(dictionary)
	fmt.Println(dictionary["first"])

	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	definition, err = dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err = dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	} else {
		definition, _ = dictionary.Search("hello")
		fmt.Println("found hello, definition:", definition)
	}

	err = dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	} else {
		definition, _ = dictionary.Search("hello")
		fmt.Println("found hello, definition:", definition)
	}
}
