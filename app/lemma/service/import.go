package service

import (
	"encoding/xml"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"parser/app/lemma/entities"
	"sync"
	"time"
)

func StartImportToChan(wg *sync.WaitGroup) <-chan entities.Lemma {
	t := time.Now()
	lemmaChan := make(chan entities.Lemma, 1000)

	wg.Add(1)

	fmt.Println(time.Now().Sub(t))

	path := viper.GetString("ROOT_PATH") + "/xml/dict.opcorpora.xml"
	fmt.Println("Запускаем горутину чтения XML в канал.")
	go readXmlToChan(path, lemmaChan, wg)

	return lemmaChan
}

func readXmlToChan(path string, lemmaChan chan entities.Lemma, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(lemmaChan)
	filePath, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	decoder := xml.NewDecoder(filePath)

	for {
		token, error := decoder.Token()

		if error != nil {
			if error == io.EOF {
				log.Println("Дочитали XML файл до конца")
			} else {
				log.Fatal(error)
			}
			break
		}

		// Типа того, что token.(type) извлекаем тип через рефлексию.
		switch el := token.(type) {
		case xml.StartElement:
			if el.Name.Local == "lemma" {
				//fmt.Println(el)
				lem := new(entities.Lemma)
				decoder.DecodeElement(lem, &el)
				//fmt.Println(lem)
				lemmaChan <- *lem
			}
		}
	}
}
