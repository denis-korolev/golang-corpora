package tests

import (
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
