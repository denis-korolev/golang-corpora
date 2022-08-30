package utils

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"parser/entities"
	"sync"
)

func CreateFile(fileName string) {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
}

func WriteDataToFile(data string, fileName string) {

	f, operror := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0644)
	if operror != nil {
		log.Fatal(operror)
	}
	str := data
	str += "\n"
	_, wrerror := f.WriteString(str)
	if wrerror != nil {
		f.Close()
		log.Fatal(wrerror)
	}
	f.Close()
}

func OpenFileFs(dir string, filename string) ([]byte, error) {
	fsys := os.DirFS(dir)
	data, error := fs.ReadFile(fsys, filename)

	if error != nil {
		return []byte{}, error
	}

	return data, error
}

func OpenFileOs() {
	file, error := os.Open("xml/OpcorporaTestingFile.xml")

	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()
	dataBuffer := make([]byte, 1)

	for {
		n, readError := file.Read(dataBuffer)

		if readError == io.EOF {
			break
		}

		if readError != nil {
			log.Fatal(readError)
		}

		if n > 0 {
			fmt.Println(string(dataBuffer))
		}
	}
}

func ReadXmlToChan(path string, lemmaChan chan entities.Lemma, wg *sync.WaitGroup) {
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
				log.Println("Дочитали до конца")
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
