package main

import (
	"log"
	"parser/entities"
	"parser/reader"
	"parser/storage"
	"runtime"
	"sync"
	"time"
)

func main() {
	t := time.Now()

	log.Println("Получаем читатель")
	xmlReader, _ := reader.NewLemmaProviderFromFile("./xml/OpcorporaTestingFile.xml")
	log.Println(time.Since(t))

	lemmaStorage := storage.CreateLemmaStorage(storage.WithLogger|storage.WithTimer, nil)

	//запустить как горутину
	wg := new(sync.WaitGroup)
	log.Println("Запускаем горутины.")
	wg.Add(1)
	lemmaChan := xmlReader.GetLemmasChan(runtime.NumCPU()*2, wg)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go sendLemmaToStorage(lemmaChan, lemmaStorage, wg)
	}

	wg.Wait()
	log.Println("All goroutines complete.")
	log.Println("Ушло времени:", time.Since(t))
}

func sendLemmaToStorage(lemmaChan <-chan *entities.Lemma, lemmaStorage storage.LemmaStorageInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	for l := range lemmaChan {
		err := lemmaStorage.InsertNewLemma(l)
		if err != nil {
			log.Println(err)
		}
	}
}
