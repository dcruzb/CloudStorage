package cloudLib

import (
	"os"
	"time"
)

type StorageFunctions interface {
	SendFile(file *os.File) (createdFile CloudFile, err error)
	GetFile() (file *os.File, err error)
	List(path string) (files []CloudFile, err error)
}

type CloudFile struct {
	Id          string
	Path        string
	Cloud       string
	Size        string
	Created     time.Time
	LastChecked time.Time
}
