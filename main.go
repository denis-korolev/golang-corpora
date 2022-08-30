package main

import (
	"context"
	"encoding/json"
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
	var wg sync.WaitGroup
	t := time.Now()
	lemmaChan := make(chan entities.Lemma, 1000)

	es, err := clients.CreateElasticClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	fmt.Println(time.Now().Sub(t))
	fmt.Println("Запускаем чтение XML в канал.")
	//"OpcorporaTestingFile.xml" "dict.opcorpora.xml"
	wg.Add(1)
	go utils.ReadXmlToChan("xml/dict.opcorpora.xml", lemmaChan, &wg)

	fmt.Println("Запускаем горутины для эластика.")
	fmt.Println(time.Now().Sub(t))
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go sendLemmaToElastic(lemmaChan, es, &wg)
	}

	wg.Wait()
	fmt.Println("All goroutines complete.")
	fmt.Println(time.Now().Sub(t))

	fmt.Println("Нажми на любую клавишу")
	var input string
	fmt.Scanln(&input)
}

func sendLemmaToElastic(lemmaChan chan entities.Lemma, es *elasticsearch.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	for m := range lemmaChan {
		jsonData, err := json.Marshal(m)
		if err != nil {
			log.Fatal("Ошибка маршалинга в json: %s", err)
		}

		// Set up the request object.
		req := esapi.IndexRequest{
			Index:      "lemma",
			DocumentID: m.ID,
			Body:       strings.NewReader(string(jsonData)),
			Refresh:    "true",
		}

		//// Perform the request with the client.
		res, err := req.Do(context.Background(), es)
		//log.Println(res)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("[%s] Error indexing document ID=%d", res.Status())
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
