package service

import (
	"parser/app/lemma/entities"
	"sync"
	"testing"
)

func TestReadXmlToChan(t *testing.T) {
	var wg sync.WaitGroup
	lemmaChan := make(chan entities.Lemma, 20)
	wg.Add(1)

	readXmlToChan("./../../../common/OpcorporaTestingFile.xml", lemmaChan, &wg)

	if len(lemmaChan) != 2 {
		t.Fail()
	}

	for lemma := range lemmaChan {
		if lemma.ID == "1" && lemma.L.T != "ёж" {
			t.Fail()
		}
		if lemma.ID == "2" && lemma.L.T != "ёжик" {
			t.Fail()
		}
	}
	wg.Wait()
}
