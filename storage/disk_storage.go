/*
author: foolbread
file: storage/disk_storage.go
date: 2017/9/27
*/
package storage

import (
	"io"
	"io/ioutil"
	"os"
)

type DiskStorage struct {

}

func (s *DiskStorage)ChangeDir(dir string)(bool,error){
	info,err := os.Stat(dir)
	if err != nil{
		return false,err
	}

	return info.IsDir(),nil
}


func (s *DiskStorage)ListFile(dir string)([]*FTPFileInfo,error){
	infos,err := ioutil.ReadDir(dir)
	if err != nil{
		return nil,err
	}

	var ret []*FTPFileInfo
	for _,v := range infos{
		f := newFTPFileInfo()
		f.Name = v.Name()
		f.IsDir = v.IsDir()
		f.Mode = v.Mode()
		f.Size = v.Size()
		f.ModTime = v.ModTime()

		ret = append(ret,f)
	}

	return ret,nil
}

func (s *DiskStorage)ReName(oldname string, newname string)error{
	return nil
}

func (s *DiskStorage)MKDir(dir string)error{
	return nil
}

func (s *DiskStorage)DeleteFile(filename string)error{
	return nil
}

func (s *DiskStorage)DeleteDir(dir string)error{
	return nil
}

func (s *DiskStorage)GetFile(filename string)(io.ReadCloser,error){
	return nil,nil
}

func (s *DiskStorage)StoreFile(filename string,rd io.Reader)(int64,error){
	return 0,nil
}
