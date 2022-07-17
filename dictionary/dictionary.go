package dictionary

//Пакет dictionary предоставляет функции и методы для чтения и подсчета повторенных слов.
//Пакет дает доступ на использование лишь 3 методов и 1 функции, так как доступ к другим методам и отдельным полям структур может помешать работе необходимых процедур.

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

//Структура dictionary хранит в себе срез структур word и ссылку на открытый файл
type dictionary struct {
	words []word
	file  *os.File
}

//Структура word хранит в себе срез байтов репрезентующих одно слово и числовое значение репрезентующее количество повторений среза
type word struct {
	word    []byte
	counter int
}

//Функция NewDictionary возвращает объект структуры dictionary с необходимыми полями.
func NewDictionary(file *os.File) *dictionary {
	words := make([]word, 0, 10000)
	dictionary := dictionary{
		words: words,
		file:  file,
	}

	return &dictionary
}

//Метод PerformSearch заполняет срез words у своего объекта.
//Логика метода описана в комментариях внутри кода метода.
func (d *dictionary) PerformSearch() {

	//Создаем буфферы для записи и чтения.
	//Буффер на чтение имеет размер в один байт, так как чтение будет осуществляться по одному символу
	readingBuffer := make([]byte, 1)
	//Буффер на запись имеет размер в 32 байта, так как этот размер позволит нам не производить дальнейшее расширение среза.
	wordBuffer := make([]byte, 0, 32)

	//P.s Здесь мог быть использован пакет Bufio, однако я посчитал, что использования обычных срезов более чем удовлетворяет потребности метода

	for {
		//Чтение файла по одному символу и запись символа в массив для чтения в виде байта
		n, err := d.file.Read(readingBuffer)

		if n > 0 {

			if readingBuffer[0] >= byte('A') && readingBuffer[0] <= byte('Z') {
				//Если символ является буквой в большом регистре, то мы записываем ее аналог в нижнем регистре в массив для записи
				readingBuffer[0] = readingBuffer[0] + byte(' ')
				wordBuffer = append(wordBuffer, readingBuffer[0])

			} else if readingBuffer[0] >= byte('a') && readingBuffer[0] <= byte('z') {
				//Если символ является буквой в нижнем регистре, то мы записываем ее в массив для записи
				wordBuffer = append(wordBuffer, readingBuffer[0])

			} else if readingBuffer[0] == byte(' ') && len(wordBuffer) != 0 {
				//Если символ является буквой пробелом и буффер для записи не пустой, то мы передаем буффер записи в метод addWord
				d.addWord(wordBuffer)
				//После очищаем контент буффера, оставляя выделенную им память.
				wordBuffer = wordBuffer[:0]

			} else if ((readingBuffer[0] > byte('z') || readingBuffer[0] < byte('A')) || (readingBuffer[0] > byte('Z') && readingBuffer[0] < byte('a'))) && len(wordBuffer) != 0 {
				//Если символ является буквой пробелом и буффер для записи не пустой, то мы передаем буффер записи в метод addWord
				d.addWord(wordBuffer)
				//После очищаем контент буффера, оставляя выделенную им память.
				wordBuffer = wordBuffer[:0]

			}

		} else if err == io.EOF {
			//Если пришли к концу файла, то мы передаем буффер записи в метод addWord
			d.addWord(wordBuffer)
			//Выходим из цикла. Очищать буффер необязательно, за нас это может сделать сборщик мусора
			break
		}
	}
}

//Метод PrintResults печатает requestedTop строк, в которых написаны слова и частота их повторений
func (d *dictionary) PrintResults(requestedTop int) {

	if requestedTop > len(d.words) {
		fmt.Println("requested amount of words is too big for pool of", len(d.words), "words")
		return
	}
	for i := 0; i < requestedTop; i++ {
		fmt.Println(strconv.Itoa(d.words[i].counter) + " " + string(d.words[i].word))
	}
}

//Метод PrintResultsToWriter печатает requestedTop строк в буфер. Этот метод вызывается в main_test.go
func (d *dictionary) PrintResultsToWriter(requestedTop int, out io.Writer) {
	if requestedTop > len(d.words) {
		fmt.Println("requested amount of words is too big for pool of", len(d.words), "words")
		return
	}
	for i := 0; i < requestedTop; i++ {
		fmt.Fprint(out, strconv.Itoa(d.words[i].counter)+" "+string(d.words[i].word)+"\n")
	}
}

//Метод addWord добавляет новое слово в случае если оно не было найдене в срезе слов,
//или увеличивает счетчик слова и вызывает метод swapleft в случае его обнаружения в срезе слов.
func (d *dictionary) addWord(wordBuffer []byte) {

	state, index := d.contains(wordBuffer)
	if state {

		(d.words)[index].counter++
		d.swapleft(index)

	} else {

		newSlice := make([]byte, len(wordBuffer))
		copy(newSlice, wordBuffer)
		words := word{newSlice, 1}
		d.words = append(d.words, words)

	}

}

//Метод contains проверяет наличие слова в срезе слов
func (d *dictionary) contains(element []byte) (bool, int) {

	for index, v := range d.words {
		if bytes.Equal(v.word, element) {

			return true, index
		}
	}
	return false, 0
}

//Метод swapleft производит частичную сортировку среза слов.
//Если слово по переданному индексу имеет больший счетчик чем слово до, то они поменяются местами и метод будет вызван рекурсивно.
func (d *dictionary) swapleft(nodeIndex int) {
	if nodeIndex == 0 {
		return
	}
	if d.words[nodeIndex].counter > d.words[nodeIndex-1].counter {
		d.words[nodeIndex], d.words[nodeIndex-1] = d.words[nodeIndex-1], d.words[nodeIndex]
		d.swapleft(nodeIndex - 1)
	}

}
