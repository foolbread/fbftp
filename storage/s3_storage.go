/*
author: foolbread
file: storage/s3_storage.go
date: 2017/9/27
*/
package storage

import (
	"io"
)

type S3Storage struct {

}

func (s *S3Storage)Stat(file string)(*FTPFileInfo,error){
	return nil,nil
}

func (s *S3Storage)ChangeDir(dir string)(bool,error){
	return true,nil
}

func (s *S3Storage)ListFile(dir string)([]*FTPFileInfo,error){
	return nil,nil
}

func (s *S3Storage)ReName(oldname string, newname string)error{
	return nil
}

func (s *S3Storage)MKDir(dir string)error{
	return nil
}

func (s *S3Storage)DeleteFile(filename string)error{
	return nil
}

func (s *S3Storage)DeleteDir(dir string)error{
	return nil
}

func (s *S3Storage)GetFile(filename string, wr io.Writer)(int64,error){
	return 0,nil
}

func (s *S3Storage)StoreFile(filename string, rd io.Reader)(int64,error){
	return 0,nil
}
