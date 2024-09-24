package filemanager

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"
	"errors"
)

type FileManager struct {
	ReadFilepath string 
	WriteFilePath string
}


func (f *FileManager) Read() ([][]string, error){
	byteSlice, err := os.ReadFile(f.ReadFilepath)

	if err != nil {
		return [][]string{}, err
	}

	csvData, _ := csv.NewReader(strings.NewReader(string(byteSlice))).ReadAll()

    contents := csvData[1:]

	return contents, nil
}

func (f *FileManager) Write(s interface{}) error {
	data, err := json.Marshal(s)

	if err != nil {
		return errors.New("failed to parse to json")
	}

	stringData := string(data)

	return os.WriteFile(f.WriteFilePath, []byte(stringData), 0764)
}