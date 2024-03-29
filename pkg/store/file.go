package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

type FileStore struct {
	FileName string
	Mock     *Mock
}
type Mock struct {
	Data []byte
	Err  error
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}
func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName, nil}
	}
	return nil
}

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)

}

func (fs *FileStore) Write(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		encodedData, _ := json.Marshal(data)
		fs.Mock.Data = encodedData
		return nil
	}
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.FileName, fileData, 0644)
	/* f, err := os.OpenFile(fs.FileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(fileData) */
	if err != nil {
		return err
	}
	return nil
}
