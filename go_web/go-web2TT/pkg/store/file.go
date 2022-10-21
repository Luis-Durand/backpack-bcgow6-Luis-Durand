package store

import (
	"encoding/json"
	"os"
)

/* import "path/filepath"
 */

type Type string

const (
	FileType Type = "file"
)

type Store interface {
	Write(data interface{}) error
	Read(data interface{}) error
}

type fileStore struct {
	FilePath string
}

func NewStore(store Type, fileName string) Store {

	switch store {
	case FileType:
		return &fileStore{fileName}
	}

	return nil

}

func (f *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(f.FilePath, fileData, 0644)
}

func (f *fileStore) Read(data interface{}) error {

	file, err := os.ReadFile(f.FilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &data)

}
