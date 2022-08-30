package tests

import (
	"parser/entities"
	"parser/utils"
	"regexp"
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

	lemmaChan := make(chan entities.Lemma, 20)
	error := utils.ReadXmlToChan("./../xml/OpcorporaTestingFile.xml", lemmaChan)

	if error != nil {
		t.Error(error)
		t.Fail()
	}

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

}
