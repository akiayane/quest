# quest
 
Пакет dictionary предоставляет функции и методы для чтения и подсчета повторенных слов.
Пример использования пакета можно увидеть в файле main.go.

Что бы запустить тест используйте "go test -v". Данные для теста взяты из результатов bash команды cat mobydick.txt | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr | head -20

Что бы запустить тест производительности используйте "go test -bench . -benchmem". Результаты теста: 
![benchmark results](https://user-images.githubusercontent.com/67632960/179423902-aa294d49-924e-45a6-9c9f-efb085d67839.png)
