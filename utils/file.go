package utils

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
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
	file, error := os.Open("common/OpcorporaTestingFile.xml")

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
