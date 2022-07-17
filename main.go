package main

import (
	"kazdream-quest/dictionary"
	"log"
	"os"
)

func main() {
	value := os.Getenv("filename")
	if len(value) == 0 {
		value = "mobydick.txt"
	}
	file, err := os.Open("mobydick.txt") //open file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	dictionary := dictionary.NewDictionary(file)
	dictionary.PerformSearch()
	dictionary.PrintResults(20)

}
