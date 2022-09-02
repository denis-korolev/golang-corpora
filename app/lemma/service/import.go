package service

import (
	"encoding/xml"
	"fmt"
	"github.com/nabbar/golib/archive/bz2"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"parser/app/lemma/entities"
	"strings"
	"sync"
	"time"
)

func DownloadArchive() string {
	fullURLFile := viper.GetString("CORPORA_FILE")

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	varPath := viper.GetString("VAR_PATH")
	varFile := varPath + "/" + fileName

	fmt.Println("Создаем пустой файл" + varFile)
	file, err := os.Create(varFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
		// Put content on file
	}
	fmt.Println("Скачиваем файл")
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	fmt.Printf("Downloaded a file %s with size %d \n", fileName, size)

	fmt.Println("Создаем новый файл словарей")
	corporaName := varPath + "/" + "dict.opcorpora.xml"
	corporaFile, err := os.Create(corporaName)
	if err != nil {
		log.Fatal(err)
	}
	defer corporaFile.Close()

	fmt.Println("Распаковываем архив" + varFile)
	copyError := bz2.GetFile(file, corporaFile)
	if copyError != nil {
		fmt.Println("Ошибка во время распаковки")
		log.Fatal(copyError)
	}
	fmt.Println("Готово")

	fmt.Println("Удаляем архив " + varFile)
	os.Remove(varFile)

	return corporaName
}

func StartImportToChan(path string, wg *sync.WaitGroup) <-chan entities.Lemma {
	t := time.Now()
	lemmaChan := make(chan entities.Lemma, 1000)

	wg.Add(1)

	fmt.Println(time.Now().Sub(t))
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
