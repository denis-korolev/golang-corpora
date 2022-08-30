package tests

import (
	"parser/entities"
	"parser/utils"
	"regexp"
	"sync"
	"testing"
)

func TestOpenFileFs(t *testing.T) {
	data, error := utils.OpenFileFs("./../xml", "OpcorporaTestingFile.xml")
	want := regexp.MustCompile("revision=\"417150\"")
	if !want.MatchString(string(data)) || error != nil {
		t.Fail()
	}
}

func TestReadXmlToChan(t *testing.T) {
	var wg sync.WaitGroup
	lemmaChan := make(chan entities.Lemma, 20)
	wg.Add(1)
	utils.ReadXmlToChan("./../xml/OpcorporaTestingFile.xml", lemmaChan, &wg)

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
