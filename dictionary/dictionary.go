package dictionary

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type dictionary struct {
	words *[]word
	file  *os.File
}

type word struct {
	word    []byte
	counter int
}

func (d *dictionary) swapleft(nodeIndex int) {
	if nodeIndex == 0 {
		return
	}
	if (*d.words)[nodeIndex].counter > (*d.words)[nodeIndex-1].counter {
		(*d.words)[nodeIndex], (*d.words)[nodeIndex-1] = (*d.words)[nodeIndex-1], (*d.words)[nodeIndex]
	}
	d.swapleft(nodeIndex - 1)
}

func (d *dictionary) contains(element []byte) (bool, int) {

	for index, v := range *d.words {
		if bytes.Equal(v.word, element) {

			return true, index
		}
	}
	return false, 0
}

func (d *dictionary) addWord(wordBuffer []byte) {

	state, index := d.contains(wordBuffer)
	if state {

		(*d.words)[index].counter++
		d.swapleft(index)

	} else {

		newSlice := make([]byte, len(wordBuffer))
		copy(newSlice, wordBuffer)
		words := word{newSlice, 1}
		*d.words = append(*d.words, words)

	}

}

func NewDictionary(file *os.File) *dictionary {
	words := make([]word, 0, 10000)
	dictionary := dictionary{
		words: &words,
		file:  file,
	}

	return &dictionary
}

func (d *dictionary) PerformSearch() {

	readingBuffer := make([]byte, 1)
	wordBuffer := make([]byte, 0, 32)

	for {
		//reading file's letters one by one

		n, err := d.file.Read(readingBuffer)

		if n > 0 {
			if readingBuffer[0] >= byte('A') && readingBuffer[0] <= byte('Z') {

				readingBuffer[0] = readingBuffer[0] + byte(' ')
				wordBuffer = append(wordBuffer, readingBuffer[0])

			} else if readingBuffer[0] >= byte('a') && readingBuffer[0] <= byte('z') {

				wordBuffer = append(wordBuffer, readingBuffer[0])

			} else if readingBuffer[0] == byte(' ') && len(wordBuffer) != 0 {

				d.addWord(wordBuffer)
				wordBuffer = wordBuffer[:0]

			} else if ((readingBuffer[0] > byte('z') || readingBuffer[0] < byte('A')) || (readingBuffer[0] > byte('Z') && readingBuffer[0] < byte('a'))) && len(wordBuffer) != 0 {

				d.addWord(wordBuffer)
				wordBuffer = wordBuffer[:0]

			}

		} else if err == io.EOF {
			d.addWord(wordBuffer)
			break
		}
	}
}

func (d *dictionary) PrintResults(requestedTop int) {
	for i := 0; i < requestedTop; i++ {
		fmt.Println(strconv.Itoa((*d.words)[i].counter) + " " + string((*d.words)[i].word))
	}
}

func (d *dictionary) PrintResultsToWriter(requestedTop int, out io.Writer) {
	for i := 0; i < requestedTop; i++ {
		fmt.Fprint(out, strconv.Itoa((*d.words)[i].counter)+" "+string((*d.words)[i].word)+"\n")
	}
}
