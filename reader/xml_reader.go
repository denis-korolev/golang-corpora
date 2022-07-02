package reader

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"parser/entities"
	"sync"
)

type LemmaProviderInterface interface {
	GetLemmasChan(bufferSize int, wg *sync.WaitGroup) <-chan *entities.Lemma
}

type XmlLemmaReader struct {
	input io.Reader
}

func NewXmlLemmaProviderFormStream(reader io.Reader) *XmlLemmaReader {
	return &XmlLemmaReader{input: reader}
}

func NewLemmaProviderFromFile(filePath string) (*XmlLemmaReader, error) {
	input, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return &XmlLemmaReader{
		input: input,
	}, nil
}

func (r *XmlLemmaReader) GetLemmasChan(bufferSize int, wg *sync.WaitGroup) <-chan *entities.Lemma {

	output := make(chan *entities.Lemma, bufferSize)

	go func() {
		defer wg.Done()

		r.decodeXml(output)

		close(output)
	}()

	return output
}

func (r *XmlLemmaReader) decodeXml(output chan *entities.Lemma) {

	decoder := xml.NewDecoder(r.input)

	for {
		t, err := decoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error decoding token: %s", err)
		} else if t == nil {
			break
		}

		switch se := t.(type) {

		case xml.StartElement:
			switch se.Name.Local {

			// Found an item, so we process it
			case "lemma":
				var item = new(entities.Lemma)

				if err = decoder.DecodeElement(item, &se); err != nil {
					log.Fatalf("Error decoding item: %s", err)
				}
				output <- item
			}
		}
	}
}
