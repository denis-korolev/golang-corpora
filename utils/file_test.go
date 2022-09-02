package utils

import (
	"regexp"
	"testing"
)

func TestOpenFileFs(t *testing.T) {
	data, error := OpenFileFs("./../xml", "OpcorporaTestingFile.xml")
	want := regexp.MustCompile("revision=\"417150\"")
	if !want.MatchString(string(data)) || error != nil {
		t.Fail()
	}
}
