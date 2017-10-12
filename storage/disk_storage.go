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

func (s *DiskStorage)Stat(file string)(*FTPFileInfo,error){
	info,err := os.Stat(file)
	if err != nil{
		return nil,err
	}

	ret := newFTPFileInfo()
	ret.Name = info.Name()
	ret.IsDir = info.IsDir()
	ret.Mode = info.Mode()
	ret.Size = info.Size()
	ret.ModTime = info.ModTime()

	return ret,nil
}

func (s *DiskStorage)ChangeDir(dir string)(bool,error){
	info,err := os.Stat(dir)
	if err != nil{
		return false,err
	}

	return info.IsDir(),nil
}


func (s *DiskStorage)ListFile(file string)([]*FTPFileInfo,error){
	info,err := os.Stat(file)
	if err != nil{
		return nil,err
	}

	var ret []*FTPFileInfo
	if info.IsDir(){
		infos,err := ioutil.ReadDir(file)
		if err != nil{
			return nil,err
		}

		for _,v := range infos{
			f := newFTPFileInfo()
			f.Name = v.Name()
			f.IsDir = v.IsDir()
			f.Mode = v.Mode()
			f.Size = v.Size()
			f.ModTime = v.ModTime()

			ret = append(ret,f)
		}
	}else{
		f := newFTPFileInfo()
		f.Name = info.Name()
		f.IsDir = info.IsDir()
		f.Mode = info.Mode()
		f.Size = info.Size()
		f.ModTime = info.ModTime()

		ret = append(ret,f)
	}

	return ret,nil
}

func (s *DiskStorage)ReName(oldname string, newname string)error{
	return nil
}

func (s *DiskStorage)MKDir(dir string)error{
	return os.Mkdir(dir,0777)
}

func (s *DiskStorage)DeleteFile(filename string)error{
	return nil
}

func (s *DiskStorage)DeleteDir(dir string)error{
	return os.Remove(dir)
}

func (s *DiskStorage)GetFile(filename string,wr io.Writer)(int64,error){
	fl,err := os.Open(filename)
	if err != nil{
		return 0,err
	}
	defer fl.Close()

	return io.Copy(wr,fl)
}

func (s *DiskStorage)StoreFile(filename string,rd io.Reader)(int64,error){
	fl,err := os.OpenFile(filename,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0777)
	if err != nil{
		return 0,err
	}
	defer fl.Close()

	return io.Copy(fl,rd)
}
