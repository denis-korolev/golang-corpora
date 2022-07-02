package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"parser/clients"
	"parser/entities"
	"parser/utils"
	"strings"
	"sync"
	"time"
)

func main() {
	var newDictionary entities.Dictionary
	var lemmaJson entities.LemmaJson

	t := time.Now()
	lemmaChan := make(chan string)

	//"OpcorporaTestingFile.xml" "dict.opcorpora.xml"
	fmt.Println("Открываем XML.")
	fileData, error := utils.OpenFileFs("xml", "dict.opcorpora.xml")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(time.Now().Sub(t))

	fmt.Println("Конвертируем XML в структуру.")
	xml.Unmarshal(fileData, &newDictionary)
	fmt.Println(time.Now().Sub(t))

	es, err := clients.CreateElasticClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	//запустить как горутину
	fmt.Println("Запускаем горутины.")
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go sendLemmaToElastic(lemmaChan, lemmaJson, es, &wg)
	}
	fmt.Println(time.Now().Sub(t))

	fmt.Println("Отправляем все в эластик.")
	for _, lemma := range newDictionary.Lemmata.Lemma {
		lemmaLoc := lemma
		m, err := json.Marshal(lemmaLoc)
		if err != nil {
			log.Fatal(err)
		}
		lemmaChan <- string(m)
	}
	close(lemmaChan)

	wg.Wait()
	fmt.Println("All goroutines complete.")
	fmt.Println(time.Now().Sub(t))

	fmt.Println("Нажми на любую клавишу")
	var input string
	fmt.Scanln(&input)
}

func sendLemmaToElastic(lemmaChan chan string, lemmaJson entities.LemmaJson, es *elasticsearch.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	for m := range lemmaChan {
		json.Unmarshal([]byte(m), &lemmaJson)
		//fmt.Println(lemmaJson)

		// Set up the request object.
		req := esapi.IndexRequest{
			Index:      "lemma",
			DocumentID: lemmaJson.ID,
			Body:       strings.NewReader(m),
			Refresh:    "true",
		}

		// Perform the request with the client.
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("[%s] Error indexing document ID=%d", res.Status(), lemmaJson.ID)
		} else {
			// Deserialize the response into a map.
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				log.Printf("Error parsing the response body: %s", err)
			} else {
				// Print the response status and indexed document version.
				//log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
			}
		}

	}
}
