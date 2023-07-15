package main

import (
	"fmt"
	"nmc/dictionary/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "First"}
	fmt.Println(dictionary)
	fmt.Println(dictionary["first"])

	// READ
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

	// CREATE
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

	// UPDATE
	err = dictionary.Update("hello", "World")
	if err != nil {
		fmt.Println(err)
	} else {
		definition, _ = dictionary.Search("hello")
		fmt.Println("updated hello, definition:", definition)
	}

	err = dictionary.Update("world", "Hello")
	if err != nil {
		fmt.Println(err)
	} else {
		definition, _ = dictionary.Search("world")
		fmt.Println("updated world, definition:", definition)
	}

	// DELETE
	dictionary.Delete("hello")
	definition, err = dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
