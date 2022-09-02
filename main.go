package main

import (
	"fmt"
	"log"
	"parser/clients"
	"parser/config"
	"parser/entities"
	"parser/utils"
	"sync"
	"time"
)

func main() {
	config.CalculatetConfig()

	var wg sync.WaitGroup
	t := time.Now()
	lemmaChan := make(chan entities.Lemma, 1000)

	es, err := clients.CreateElasticClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	fmt.Println(time.Now().Sub(t))
	fmt.Println("Запускаем чтение XML в канал.")

	wg.Add(1)
	go utils.ReadXmlToChan("xml/dict.opcorpora.xml", lemmaChan, &wg)

	fmt.Println("Запускаем горутины для эластика.")
	fmt.Println(time.Now().Sub(t))

	clients.BulkLemma(lemmaChan, es)

	wg.Wait()
	fmt.Println("All goroutines complete.")
	fmt.Println(time.Now().Sub(t))

	fmt.Println("Нажми на любую клавишу")
	var input string
	fmt.Scanln(&input)
}
