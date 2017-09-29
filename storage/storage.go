/*
author: foolbread
file: storage/storage.go
date: 2017/8/10
*/
package storage

import (
	"github.com/foolbread/fbcommon/golog"
	"time"
	"io"
	"os"
)

func InitStorage(){
	golog.Info("fbftp storage initing......")
}

const(
	default_owner = "999"
	default_group = "999"
)

type FTPFileInfo struct {
	Name string
	IsDir bool
	Mode os.FileMode
	Size int64
	ModTime time.Time
	Owner string
	Group string
}

func newFTPFileInfo()*FTPFileInfo{
	r := new(FTPFileInfo)
	r.Owner = default_owner
	r.Group = default_group

	return r
}

func (s *FTPFileInfo)TransforModeToString()string{
	return ""
}

type FTPStorage interface {
	ChangeDir(string)(bool,error)
	ListFile(string)([]*FTPFileInfo,error)
	ReName(string, string)error
	MKDir(string)error
	DeleteFile(string)error
	DeleteDir(string)error
	GetFile(string)(io.ReadCloser,error)
	StoreFile(string,io.Reader)(int64,error)
}