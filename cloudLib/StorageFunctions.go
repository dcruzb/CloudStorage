package cloudLib

import (
	"os"
	"time"
)

type StorageFunctions interface {
	SendFile(base64File string, fileName string, remotePath string) (createdFile CloudFile, err error)
	GetFile(fileName string, path string) (file *os.File, err error)
	List(path string) (files []CloudFile, err error)
}

type CloudFile struct {
	Id          string
	Path        string
	Cloud       string
	Size        string // Size in mb
	Created     time.Time
	LastChecked time.Time
}
