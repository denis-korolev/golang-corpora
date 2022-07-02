package main

import (
	"encoding/xml"
	"log"
	"os"
	"parser/entities"
	"parser/storage"
	"runtime"
	"sync"
	"time"
)

func main() {
	var newDictionary entities.Dictionary

	t := time.Now()
	lemmaChan := make(chan entities.Lemma, runtime.NumCPU()*2)

	//"OpcorporaTestingFile.xml" "dict.opcorpora.xml"
	log.Println("Открываем XML.")
	fileStream, err := os.Open("./xml/OpcorporaTestingFile.xml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(time.Since(t))

	t = time.Now()
	log.Println("Конвертируем XML в структуру.")
	xmlDecoder := xml.NewDecoder(fileStream)
	_ = xmlDecoder.Decode(&newDictionary)
	log.Println(time.Since(t))

	lemmaStorage := storage.CreateLemmaStorage(storage.WithLogger|storage.WithTimer, nil)

	//запустить как горутину
	t = time.Now()
	log.Println("Запускаем горутины.")
	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go sendLemmaToStorage(lemmaChan, lemmaStorage, &wg)
	}
	log.Println(time.Since(t))

	t = time.Now()
	log.Println("Отправляем все в эластик.")
	for _, lemma := range newDictionary.Lemmata.Lemma {
		lemmaChan <- lemma
	}
	close(lemmaChan)

	wg.Wait()
	log.Println("All goroutines complete.")
	log.Println(time.Since(t))
}

func sendLemmaToStorage(lemmaChan <-chan entities.Lemma, lemmaStorage storage.LemmaStorageInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	for l := range lemmaChan {
		err := lemmaStorage.InsertNewLemma(l)
		if err != nil {
			log.Println(err)
		}
	}
}
