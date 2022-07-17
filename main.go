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
	file, err := os.Open("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//Инициализируем новый объект из пакета dictionary.
	dictionary := dictionary.NewDictionary(file)
	//Сохраняем уникальные слова и количество их повторений, попутно сортируя их порядок.
	//Логика сохранение слов и подсчета их повторений описана в аннотации и комментариях пакета dictionary.
	dictionary.PerformSearch()
	//Печатаем n количество слов. N не может быть больше длины среза в котором хранятся все слова.
	dictionary.PrintResults(20)

}
