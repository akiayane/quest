package main

import (
	"bytes"
	dictionary "kazdream-quest/dictionary"
	"log"
	"os"
	"testing"
)

// -----
// go test -v

func TestSearch(t *testing.T) {
	expectedOutputBuffer := new(bytes.Buffer)
	expectedOutput := []byte("4284 the\n2192 and\n2185 of\n1861 a\n1685 to\n1366 in\n1056 i\n1024 that\n889 his\n821 it\n783 he\n616 but\n603 was\n595 with\n577 s\n564 is\n551 for\n542 all\n541 as\n458 at\n")
	expectedOutputBuffer.Write(expectedOutput)

	file, err := os.Open("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	realOutputBuffer := new(bytes.Buffer)

	dictionary := dictionary.NewDictionary(file)
	dictionary.PerformSearch()
	dictionary.PrintResultsToWriter(20, realOutputBuffer)

	if !bytes.Equal(expectedOutputBuffer.Bytes(), realOutputBuffer.Bytes()) {
		t.Errorf("results not match\nGot:\n%vExpected:\n%v", realOutputBuffer, expectedOutputBuffer)
	}
}

// -----
// go test -bench . -benchmem

func BenchmarkMt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("mobydick.txt")
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
}
